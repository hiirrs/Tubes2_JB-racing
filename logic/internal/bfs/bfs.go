package bfs

import (
	"context"
	"fmt"

	"logic/internal/entities"
	"logic/internal/tools"
)

func BFS(ctx context.Context, siteURL, targetURL string) (bool) {
	queue := []*entities.Node{}

	root := &entities.Node{
		URL: siteURL,
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

		fmt.Println("Current Node: ", currentNode.URL)

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

		childNodes, err := scraping.GetWikiNodes(ctx, currentNode.URL)
		if err != nil {
			fmt.Println(err)
			return false
		}

		currentNode.Children = append(currentNode.Children, childNodes.Children...)

		// entities.PrintTree(currentNode, 0)

		queue = append(queue, childNodes.Children...)
	}

	return false
}