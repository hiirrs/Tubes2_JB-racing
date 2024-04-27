package getPath

import (
	"context"
	"sync"
	"sync/atomic"

	"logic/internal/entities"
	scraping "logic/internal/tools"
)

func BFS(rootURL string, targetURL string, visited map[string]bool) *entities.Node {
	if rootURL == targetURL {
		return &entities.Node{URL: rootURL}
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	queue := []*entities.Node{}
	// visited := make(map[string]bool)
	var wg sync.WaitGroup
	var foundFlag int32 = 0 // Atomic flag to ensure only one goroutine sends the result

	mutex := sync.Mutex{}
	resultCh := make(chan *entities.Node, 1) // Buffer can be 1

	root := &entities.Node{
		URL:      rootURL,
		Parent:   nil,
		Children: []*entities.Node{},
		Depth:    0,
	}

	// fmt.Println("Scraping root node:", rootURL)
	scraping.ScrapeToNode(root, root.Depth)
	queue = append(queue, root.Children...)
	visited[root.URL] = true

	concurrency := len(root.Children)
	if concurrency > 1500 {
		concurrency = 1500
	}

	for len(queue) > 0 && atomic.LoadInt32(&foundFlag) == 0 {

		for i := 0; i < concurrency; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				mutex.Lock()
				currentNode := queue[i]
				mutex.Unlock()

				if visited[currentNode.URL] {
					return
				}

				if currentNode.URL == targetURL && atomic.CompareAndSwapInt32(&foundFlag, 0, 1) {
					// fmt.Println("Target found:", currentNode.URL)
					resultCh <- currentNode
					cancel() // Cancel all goroutines
					return
				}

				scraping.ScrapeToNode(currentNode, currentNode.Depth)

				mutex.Lock()
				if !visited[currentNode.URL] {
					visited[currentNode.URL] = true
					for _, child := range currentNode.Children {
						if !visited[child.URL] {
							queue = append(queue, child)
						}
					}
				}
				mutex.Unlock()
			}(i)
		}

		wg.Wait()
		if concurrency < 1000 {
			if len(queue) < 1000 {
				concurrency = len(queue)
			} else {
				concurrency = 1000
			}
		}
		queue = queue[concurrency:] // Update queue outside the goroutines
	}

	close(resultCh)
	select {
	case result := <-resultCh:
		return result
	default:
		return nil
	}
}
