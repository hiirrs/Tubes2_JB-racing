package getPath

import (
	"context"
	"fmt"
	"logic/internal/entities"
	"logic/internal/tools"
	// "sync"
	// "sync/atomic"
)

func IDS(rootURL string, targetURL string, maxDepth int) *entities.Node {
	if rootURL == targetURL {
		return &entities.Node{URL: rootURL}
	}

	root := &entities.Node{
		URL:      rootURL,
		Parent:   nil,
		Children: []*entities.Node{},
		Depth:    0,
	}

	var found *entities.Node
	for iteration := 1; iteration <= maxDepth && found == nil; iteration++ {
		found = depthLimitedSearch(root, targetURL, iteration)
	}

	return found
}

func depthLimitedSearch(root *entities.Node, targetURL string, depth int) *entities.Node {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// var wg sync.WaitGroup
	visited := make(map[string]bool)
	stack := []*entities.Node{root}
	visited[root.URL] = true

	for len(stack) > 0 {
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop from stack

		fmt.Println("Scraping node:", currentNode.URL)
		fmt.Println("Depth:", currentNode.Depth)
		fmt.Println("Visited nodes:", len(visited))

		if currentNode.Depth == depth-1 {
			fmt.Println("Check")
			scraping.ScrapeToNode(currentNode, currentNode.Depth)
			for _, child := range currentNode.Children {
				if !visited[child.URL] {
					visited[child.URL] = true
					if child.URL == targetURL {
						fmt.Println("Target found:", child.URL)
						cancel()
						return child
					}
				}
			}
		} else if currentNode.Depth < depth {
			scraping.ScrapeToNode(currentNode, currentNode.Depth)
			for _, child := range currentNode.Children {
				if !visited[child.URL] {
					visited[child.URL] = true
					child.Parent = currentNode
					stack = append(stack, child)
				}
			}
		} 
	}

	return nil
}
