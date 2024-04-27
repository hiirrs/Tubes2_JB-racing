package getPath

import (
	"context"
	"fmt"
	"logic/internal/entities"
	scraping "logic/internal/tools"
	// "sync"
	// "sync/atomic"
)

func IDS(rootURL string, targetURL string, maxDepth int) (*entities.Node, int) {
	if rootURL == targetURL {
		return &entities.Node{URL: rootURL}, 1
	}

	root := &entities.Node{
		URL:      rootURL,
		Parent:   nil,
		Children: []*entities.Node{},
		Depth:    0,
	}

	visitCount := 0

	var found *entities.Node
	var temp int
	for iteration := 1; iteration <= maxDepth && found == nil; iteration++ {
		found, temp = depthLimitedSearch(root, targetURL, iteration)
		visitCount += temp
	}

	return found, visitCount
}

func depthLimitedSearch(root *entities.Node, targetURL string, depth int) (*entities.Node, int) {
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
						return child, len(visited)
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

	return nil, len(visited)
}

func Backtrack(node *entities.Node, path []string) []string {
	if node.Parent == nil {
		// println("terakhir")
		// println(node.URL)
		path = append(path, node.URL)
		return path
	}
	// println(node.URL)
	path = append(path, node.URL)
	// println(path)
	nodes := node.Parent
	return Backtrack(nodes, path)
}

func ReverseArray(arr []string) []string {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}
