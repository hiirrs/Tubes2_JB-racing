package scraping

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
	"strings"
	"regexp"

	"logic/internal/entities"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
)



// GetPage call the client page by HTTP request and extract the body to HTML document.
func GetPage(ctx context.Context, method, siteURL string, timeout int) (*goquery.Document, error) {
	// This function can handle both all methods.
	// Initiate this body variable as nil for method that doesn't required body.
	body := io.Reader(nil)
	// fmt.Println("masuk GetPage")

	// Create a new HTTP request with context.
	req, err := http.NewRequestWithContext(ctx, method, siteURL, body)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("failed to create http request context: %w", err)
	}
	// fmt.Println("masuk request created")


	// Use default http Client.
	httpClient := &http.Client{
		Transport:     http.DefaultTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Duration(timeout) * time.Millisecond,
	}

	// Execute the request.
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("failed to execute http request: %w", err)
	}

	// fmt.Println("masuk request executed")

	// Close the response body
	defer func() { _ = resp.Body.Close() }()
	// // Parsing response body to HTML document reader.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("failed to parse html: %w", err)
	}

	// fmt.Println("masuk doc.goquery created")
	// Return HTML doc, cookies.
	return doc, nil
}

func GetLanguageCode(url string) string {
    parts := strings.Split(url, ".")
    if len(parts) >= 2 {
        return parts[0][8:] // Extract the language code after "https://"
    }
    return ""
}

func isWikipediaArticle(url string, languageCode string) bool {
    pattern := fmt.Sprintf(`^https:\/\/%s\.wikipedia\.org\/wiki\/[^:]+$`, languageCode)
    match, _ := regexp.MatchString(pattern, url)

    return match
}

// See also wikis page
func GetWikiNodes(ctx context.Context, siteURL string, languageCode string) (nodes *entities.Node, err error){
	// fmt.Println("masuk GetWikiNodes")
	doc, err := GetPage(ctx,  http.MethodGet, siteURL, 300000)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// fmt.Println("masuk GetWikiNodes ga error")


	nodes = &entities.Node{
		URL: siteURL,
	}

	doc.Find(".mw-content-ltr.mw-parser-output ul").Each(func(i int, ul *goquery.Selection){
		// fmt.Printf("masuk find ul")
		ul.Find("li").Each(func(j int, li *goquery.Selection){
			// fmt.Printf("masuk find li")

			childSiteURL, exist := li.Find("a[href*='/wiki/']").Attr("href")
			if exist {
				// fmt.Printf("URL: %s\n", childSiteURL)
				childSiteURL = adjustLink(childSiteURL)
				if isWikipediaArticle(childSiteURL, languageCode) {
					nodes.AddChild(&entities.Node{URL: childSiteURL,})
				}
			}
		})
	})

	return nodes, nil
}