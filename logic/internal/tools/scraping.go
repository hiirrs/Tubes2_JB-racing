package scraping

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"logic/internal/entities"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
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

	if !match {
		return false
	} else {
		return !strings.Contains(url, "(identifier)")
	}
}

func GetChildLinks(url string, languageCode string) (childLinks []string) {
	// Create a new Collector
	c := colly.NewCollector(
		// Restrict crawling to only wikipedia.org domain
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// Slice to store the links found
	// var childLinks []string

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)

		// Check if the link starts with '/wiki/' (Wikipedia article)
		if strings.HasPrefix(link, "/wiki/") && isWikipediaArticle(link, languageCode) {
			// Add the link to the slice of links
			childLinks = append(childLinks, link)

		}
	})

	// Start scraping on the provided URL
	c.Visit(url)

	return childLinks
}

// See also wikis page
func GetWikiNodes(ctx context.Context, siteURL string, languageCode string) (nodes *entities.Node, err error) {
	// fmt.Println("masuk GetWikiNodes")
	doc, err := GetPage(ctx, http.MethodGet, siteURL, 300000)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// fmt.Println("masuk GetWikiNodes ga error")

	nodes = &entities.Node{
		URL: siteURL,
	}

	doc.Find(".mw-content-ltr.mw-parser-output ul").Each(func(i int, ul *goquery.Selection) {
		// fmt.Printf("masuk find ul")
		ul.Find("li").Each(func(j int, li *goquery.Selection) {
			// fmt.Printf("masuk find li")

			childSiteURL, exist := li.Find("a[href*='/wiki/']").Attr("href")
			if exist {
				// fmt.Printf("URL: %s\n", childSiteURL)
				childSiteURL = adjustLink(childSiteURL)
				if isWikipediaArticle(childSiteURL, languageCode) {
					nodes.AddChild(&entities.Node{URL: childSiteURL})
				}
			}
		})
	})

	return nodes, nil
}

type Links map[string][]string

func GetAllWikiLinks(ctx context.Context, languageCode string, queueLinks []string) chan Links {
	var wg sync.WaitGroup
	res := make(chan Links)

	go func() {
		defer close(res)

		var mu sync.Mutex
		links := make(Links)

		for _, link := range queueLinks {
			wg.Add(1)
			childLinks, err := GetWikiLinks(ctx, link, languageCode, &wg)
			if err != nil {
				// Handle error
				log.Error(err)
				continue // Continue to the next link in case of error
			}
			// Merge the childLinks into the 'links' map
			mu.Lock()
			links[link] = childLinks
			mu.Unlock()
		}
		wg.Wait()
		res <- links
	}()
	return res
}

func GetWikiLinks(ctx context.Context, siteURL, languageCode string, wg *sync.WaitGroup) (childLinks []string, err error) {
	defer wg.Done()
	// fmt.Println("masuk GetWikiNodes")
	doc, err := GetPage(ctx, http.MethodGet, siteURL, 300000)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	// fmt.Println("masuk GetWikiNodes ga error")
	childs := make([]string, 0)
	doc.Find(".mw-content-ltr.mw-parser-output ul").Each(func(i int, ul *goquery.Selection) {
		// fmt.Printf("masuk find ul")
		ul.Find("li").Each(func(j int, li *goquery.Selection) {
			// fmt.Printf("masuk find li")

			childSiteURL, exist := li.Find("a[href*='/wiki/']").Attr("href")
			if exist {
				// fmt.Printf("URL: %s\n", childSiteURL)
				childSiteURL = adjustLink(childSiteURL)
				if isWikipediaArticle(childSiteURL, languageCode) {
					childs = append(childs, childSiteURL)
				}
			}
		})
	})
	return childs, nil
}

func PrintLinks(links Links) {
	for parentURL, childURLs := range links {
		fmt.Printf("Parent URL: %s\n", parentURL)
		fmt.Println("Child URLs:")
		for _, childURL := range childURLs {
			fmt.Printf(" - %s\n", childURL)
		}
		fmt.Println()
	}
}

func GetWikiLink(ctx context.Context, siteURL, languageCode string) (childLinks []string, err error) {
	// fmt.Println("masuk GetWikiNodes")
	doc, err := GetPage(ctx, http.MethodGet, siteURL, 300000)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	// fmt.Println("masuk GetWikiNodes ga error")
	childs := make([]string, 0)
	doc.Find(".mw-content-ltr.mw-parser-output ul").Each(func(i int, ul *goquery.Selection) {
		// fmt.Printf("masuk find ul")
		ul.Find("li").Each(func(j int, li *goquery.Selection) {
			// fmt.Printf("masuk find li")

			childSiteURL, exist := li.Find("a[href*='/wiki/']").Attr("href")
			if exist {
				// fmt.Printf("URL: %s\n", childSiteURL)
				childSiteURL = adjustLink(childSiteURL)
				if isWikipediaArticle(childSiteURL, languageCode) {
					childs = append(childs, childSiteURL)
				}
			}
		})
	})
	return childs, nil
}
