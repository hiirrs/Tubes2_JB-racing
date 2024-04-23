package getPath

import (
	"context"
	"fmt"

	"logic/internal/entities"
	"logic/internal/tools"
)

func BFSUpdate(ctx context.Context, siteURL, targetURL string, languageCode string) bool {
    queue := []string{} // Queue of URLs to visit
    visited := make(map[string]bool) // Map to keep track of visited URLs

    queue = append(queue, siteURL) // Add the starting URL to the queue

    for len(queue) > 0 {
        // Dequeue the URL at the front of the queue
        currentURL := queue[0]
        queue = queue[1:]

        // If the current URL is the target URL, return true
        if currentURL == targetURL {
            return true
        }

        // Skip if already visited
        if visited[currentURL] {
            continue
        }

        // Mark current URL as visited
        visited[currentURL] = true

        // Get the child URLs of the current URL
        links := <- scraping.GetAllWikiLinks(ctx, languageCode, []string{currentURL})
		// scraping.PrintLinks(links)

        // Enqueue the child URLs
		queue = append(queue, links[currentURL]...)
    }

    // If target URL not found, return false
    return false
}


func BFS(ctx context.Context, siteURL, targetURL string, languageCode string) (bool) {
	queue := []*entities.Node{}

	root := &entities.Node{
		URL: siteURL,
		Depth: 0,
	}

	queue = append(queue, root)

	visited := make(map[string]bool)

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		if visited[currentNode.URL] {
			continue
		}

		visited[currentNode.URL] = true

		// fmt.Println("Current Node: ", currentNode.URL)
		// fmt.Println("Depth: ", currentNode.Depth)

		if currentNode.URL == targetURL {
			return true
		}

		// fmt.Println("Queue: ")
		// fmt.Print("[")
		// for i := 0; i < len(queue); i++ {
		// 	fmt.Print(queue[i].URL)
		// 	if i != len(queue) - 1 {
		// 		fmt.Print(", ")
		// 	}
		// }
		// fmt.Println("]")
		// fmt.Println("Visited: ")
		// fmt.Println(visited)

		childNodes, err := scraping.GetWikiNodes(ctx, currentNode.URL, languageCode)
		if err != nil {
			fmt.Println(err)
			return false
		}

		for _, child := range childNodes.Children {
			child.Depth = currentNode.Depth + 1
			currentNode.AddChild(child)
			queue = append(queue, child)
		}
	}

	return false
}