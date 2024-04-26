package getPath

import (
	"context"
	"fmt"
	// "fmt"
	// "sync"

	"logic/internal/entities"
	"logic/internal/tools"
)

func BFS(ctx context.Context, rootURL string, targetURL string) ( *entities.Node) {
	if rootURL == targetURL {
		return &entities.Node{URL: rootURL}
	}
	
	queue := []*entities.Node{}
	visited := make(map[string]bool)

	root := &entities.Node{
		URL:      rootURL,
		Parent:   nil,
		Children: []*entities.Node{},
		Depth:    0,
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		fmt.Println("Current node: ", queue[0].URL)
		currentNode := queue[0] // Dequeue
		queue = queue[1:]

		if currentNode.URL == targetURL {
			return currentNode
		}

		if visited[currentNode.URL] {
			continue
		}

		visited[currentNode.URL] = true

		scraping.ScrapeToNode(currentNode, currentNode.Depth)

		for _, child := range currentNode.Children {
			if !visited[child.URL] {
				queue = append(queue, child) // Enqueue
			}
		}
	}

	return nil
}