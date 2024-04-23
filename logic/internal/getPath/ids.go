package getPath

import (
	"fmt"
	"sync"
	"context"

	"logic/internal/entities"
	"logic/internal/tools"
)
type PageTree struct {
	tree entities.Node
}

func SearchIDS(from, to, languageCode string, ctx context.Context) (path []string) {
	pt := PageTree{
		tree: entities.Node{
			URL: from,
		},
	}

	depth := 0
	var pathFound bool
	for (!pathFound) {
		fmt.Println("Searching depth: ", depth)
		tempPath := pt.SearchDLS(&pt.tree, to, depth, ctx, languageCode, []string{})
		if tempPath != nil {
			path = tempPath
			pathFound = true
		} else {
			depth++
		}
	}
	return path
}


func (pt *PageTree) SearchDLS(node *entities.Node,to string, depthLimit int, ctx context.Context, languageCode string, tempPath []string) []string {

	// Basis: Target page found
	if node.URL == to {
		return append(tempPath, node.URL)
	}

	// If not reach basis yet,
	// Get node children
	var err error
	if !node.Visited {
		node, err =  scraping.GetWikiNodes(ctx, node.URL, languageCode)
		node.Visited = true
		if err!=nil {
			return nil
		}
		fmt.Println("Successfully scrape ", node.URL)
	}

	// entities.PrintTree(node, 0)

	// Basis: Target page not found
	if depthLimit<=0 {
		return nil
	}

	var wg sync.WaitGroup
	ch := make(chan []string)
	for _, child := range node.Children {
		fmt.Println("Searching child: ", child.URL)

		wg.Add(1)
		go func(child *entities.Node) {
			defer wg.Done()
			if newPath := pt.SearchDLS(child, to, depthLimit-1, ctx, languageCode, append(tempPath, node.URL)); newPath != nil {
				ch <- newPath
			}
		}(child)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for newPath := range ch {
		if newPath != nil {
			return newPath
		}
	}

	return nil
}