package scraping

import (
	"regexp"
	"strings"

	"logic/internal/entities"

	"github.com/gocolly/colly/v2"
)

func ScrapeToNode (url string) *entities.Node {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	node := &entities.Node{
		URL:     url,
		Parent:  nil,
		Children: []*entities.Node{},
		Depth:   0,
	}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if isWikipediaArticle(link) {
			childNode := &entities.Node{
				URL:     "https://en.wikipedia.org" + link,
				Parent:  nil, 
				Children: []*entities.Node{},
				Depth:   0, 
			}
			node.AddChild(childNode)
		}
	})
	c.OnRequest(func(r *colly.Request) {})
	c.Visit(node.URL)
	c.Wait()

	return node
}

func isWikipediaArticle(url string) bool {
    pattern := `^\/wiki\/[^:#]+$`
    match, _ := regexp.MatchString(pattern, url)

    if !match {
        return false
    }
	
    if strings.Contains(url, "(identifier)") || strings.Contains(url, "Main_Page") {
        return false
    }

    return true
}