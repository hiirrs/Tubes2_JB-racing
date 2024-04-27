package scraping

import (
	"sync"
	"regexp"
	"strings"
	"fmt"
	"time"

	"logic/internal/entities"

	"github.com/gocolly/colly/v2"
)

func ScrapeToNode(node *entities.Node, depth int) {
    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true),
    )

	c.SetRequestTimeout(40 * time.Second)

	wg := sync.WaitGroup{}

    // Create a buffered channel to communicate between goroutines and the main routine
    childNodeChannel := make(chan *entities.Node, 10000) // Use a buffer size based on the expected number of child nodes

    // Define a function to process each HTML element concurrently
    processElement := func(e *colly.HTMLElement) {
        defer wg.Done() // Ensure wg.Done() is called regardless of path taken
        link := e.Attr("href")

        if IsWikipediaArticle(link) && ("https://en.wikipedia.org" + link) != node.URL {
            childNode := &entities.Node{
                URL:      "https://en.wikipedia.org" + link,
                Parent:   node,
                Children: []*entities.Node{},
                Depth:    node.Depth + 1,
            }
            // fmt.Println("Sending child node to channel:", childNode.URL)
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
			// fmt.Println("Receiving child node from channel:", childNode.URL)
            node.AddChild(childNode)
        }
		// fmt.Println("All child nodes received, exiting channel receiver goroutine")
    }()

    // Visit the initial URL
    err := c.Visit(node.URL)
	if err != nil {
		fmt.Println("Failed to visit URL:", node.URL, "Error:", err)
		close(childNodeChannel) // Close the channel on failure to avoid deadlock
		return
	}

    // Wait for the collector and all HTML processing to finish
    c.Wait()
	wg.Wait()

    // Close the channel after all processing is complete
	// fmt.Println("Closing child node channel")
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