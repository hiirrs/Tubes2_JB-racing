package scraping

import (
	"regexp"
	"strings"
	"sync"

	"logic/internal/entities"

	"github.com/gocolly/colly/v2"
)

func ScrapeToNode(node *entities.Node, depth int) {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	wg := sync.WaitGroup{}

	// Create a buffered channel to communicate between goroutines and the main routine
	childNodeChannel := make(chan *entities.Node, 1000000) // Use a buffer size based on the expected number of child nodes

	// Define a function to process each HTML element concurrently
	processElement := func(e *colly.HTMLElement) {
		defer wg.Done()
		link := e.Attr("href")

		if IsWikipediaArticle(link) && ("https://en.wikipedia.org"+link) != node.URL {
			childNode := &entities.Node{
				URL:      "https://en.wikipedia.org" + link,
				Parent:   node,
				Children: []*entities.Node{},
				Depth:    node.Depth + 1,
			}
			// Send the processed child node to the channel
			childNodeChannel <- childNode
		}
	}

	// Set up the parallel processing for each HTML element
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		wg.Add(1)
		go processElement(e) // Run processing in a goroutine
	})

	// Start a separate goroutine to receive processed child nodes and add them to the parent node
	go func() {
		for childNode := range childNodeChannel {
			node.AddChild(childNode)
		}
	}()

	// Visit the initial URL
	c.Visit(node.URL)

	// Wait for the collector to finish
	c.Wait()
	wg.Wait()

	// Close the channel after all processing is complete
	close(childNodeChannel)
}

func IsWikipediaArticle(url string) bool {
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
