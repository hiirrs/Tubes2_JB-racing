package getPath

import (
	"context"
	"fmt"
	"strings"

	"logic/internal/entities"
	scraping "logic/internal/tools"
)

func NewNode(URL string, depth int) *entities.Node {
	return &entities.Node{
		URL:      URL,
		Children: []*entities.Node{},
		Depth:    depth,
		Visited:  false,
	}
}

func PrintPath(path []*entities.Node) {
	var sb strings.Builder

	for i, node := range path {
		sb.WriteString(node.URL)
		if i < len(path)-1 {
			sb.WriteString(" -> ")
		}
	}

	fmt.Println("Path:", sb.String())
}

func SearchIDSC(start string, target string, languageCode string, maxDepth int) []*entities.Node {
	for depth := 0; depth <= maxDepth; depth++ {
		result, found := depthLimitedSearchC(start, target, languageCode, depth)
		if found {
			return result
		}
	}
	return nil
}

func depthLimitedSearchC(start string, target string, languageCode string, depth int) ([]*entities.Node, bool) {
	root := NewNode(start, 0)
	stack := []*entities.Node{root}
	visited := make(map[string]bool)
	paths := map[string][]*entities.Node{start: {root}}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Printf("Visiting node %s at depth %d\n", node.URL, node.Depth)

		if node.URL == target {
			return paths[node.URL], true
		}

		if node.Depth < depth && !visited[node.URL] {
			// childLinks, err := scraping.GetWikiLink(ctx, node.URL, languageCode)
			// if err != nil {
			// 	fmt.Printf("Error fetching links for node %s: %v\n", node.URL, err)
			// 	continue
			// }
			childLinks := scraping.GetChildLinks(node.URL, languageCode)

			for _, link := range childLinks {
				childNode := NewNode(link, node.Depth+1)
				node.AddChild(childNode)
				stack = append(stack, childNode)
				paths[link] = append(paths[node.URL], childNode)
			}
			visited[node.URL] = true
		}
	}
	return nil, false
}

func SearchIDS(start string, target string, ctx context.Context, languageCode string, maxDepth int) []*entities.Node {
	for depth := 0; depth <= maxDepth; depth++ {
		result := depthLimitedSearch(start, target, ctx, languageCode, depth)
		if result != nil {
			return result
		}
	}
	return nil
}

func depthLimitedSearch(start string, target string, ctx context.Context, languageCode string, depth int) []*entities.Node {
	root := NewNode(start, 0)
	stack := []*entities.Node{root}
	visited := make(map[string]bool)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Printf("Visiting node %s at depth %d\n", node.URL, node.Depth)

		if node.URL == target {
			return []*entities.Node{node}
		}

		if node.Depth < depth && !visited[node.URL] {
			childLinks, err := scraping.GetWikiLink(ctx, node.URL, languageCode)
			if err != nil {
				fmt.Printf("Error fetching links for node %s: %v\n", node.URL, err)
				continue
			}

			for _, link := range childLinks {
				childNode := NewNode(link, node.Depth+1)
				node.AddChild(childNode)
				stack = append(stack, childNode)
			}

			visited[node.URL] = true
		}
	}

	return nil
}

// func SearchIDS(start string, target string, ctx context.Context, languageCode string, maxDepth int) []*entities.Node {
// 	for depth := 0; depth <= maxDepth; depth++ {
// 		result, found := depthLimitedSearch(start, target, ctx, languageCode, depth)
// 		if found {
// 			// Backtrack to find the shortest path
// 			shortestPath := backtrack(result)
// 			return shortestPath
// 		}
// 	}
// 	return nil
// }

// func depthLimitedSearch(start string, target string, ctx context.Context, languageCode string, depth int) ([]*entities.Node, bool) {
// 	root := NewNode(start, 0)
// 	stack := []*entities.Node{root}
// 	visited := make(map[string]bool)
// 	paths := map[string][]*entities.Node{start: {root}}

// 	for len(stack) > 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		fmt.Printf("Visiting node %s at depth %d\n", node.URL, node.Depth)

// 		if node.URL == target {
// 			return paths[node.URL], true
// 		}

// 		if node.Depth < depth && !visited[node.URL] {
// 			childLinks, err := scraping.GetWikiLink(ctx, node.URL, languageCode)
// 			if err != nil {
// 				fmt.Printf("Error fetching links for node %s: %v\n", node.URL, err)
// 				continue
// 			}

// 			for _, link := range childLinks {
// 				childNode := NewNode(link, node.Depth+1)
// 				node.AddChild(childNode)
// 				stack = append(stack, childNode)
// 				paths[link] = append(paths[node.URL], childNode)
// 			}

// 			visited[node.URL] = true
// 		}
// 	}

// 	return nil, false
// }

// func backtrack(path []*entities.Node) []*entities.Node {
// 	shortestPath := make([]*entities.Node, len(path))
// 	copy(shortestPath, path)

// 	for i, j := 0, len(shortestPath)-1; i < j; i, j = i+1, j-1 {
// 		shortestPath[i], shortestPath[j] = shortestPath[j], shortestPath[i]
// 	}

// 	return shortestPath
// }

// func SearchIDS(start string, target string, ctx context.Context, languageCode string) []*entities.Node {
// 	root := NewNode(start, 0)
// 	stack := []*entities.Node{root}
// 	visited := make(map[string]bool)

// 	for len(stack) > 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		fmt.Printf("Visiting node %s at depth %d\n", node.URL, node.Depth)

// 		if node.URL == target {
// 			return []*entities.Node{node}
// 		}

// 		if !visited[node.URL] {
// 			childLinks, err := scraping.GetWikiLink(ctx, node.URL, languageCode)
// 			if err != nil {
// 				fmt.Printf("Error fetching links for node %s: %v\n", node.URL, err)
// 				continue
// 			}

// 			for _, link := range childLinks {
// 				childNode := NewNode(link, node.Depth+1)
// 				node.AddChild(childNode)
// 				stack = append(stack, childNode)
// 			}

// 			visited[node.URL] = true
// 		}
// 	}

// 	return nil
// }

// func dfssimple(node *entities.Node, target int, depth int) []int {
// 	fmt.Printf("Visiting node %d at depth %d\n", node.data, depth)
// 	if depth == 0 && node.data != target {
// 		return nil
// 	}
// 	if node.data == target {
// 		return []int{node.data}
// 	}
// 	if depth > 0 {
// 		for _, child := range node.children {
// 			result := dfs(child, target, depth-1)
// 			if result != nil {
// 				return append([]int{node.data}, result...)
// 			}
// 		}
// 	}
// 	return nil
// }

// func dfs(node *entities.Node, target string, depth int) []*entities.Node {
// 	fmt.Printf("Visiting node %d at depth %d\n", node.URL, node.Depth)
// 	if depth == 0 && node.URL != target {
// 		return nil
// 	}
// 	if node.URL == target {
// 		return []*entities.Node{node}
// 	}
// 	if depth > 0 {
// 		for _, child := range node.Children {
// 			result := dfs(child, target, depth-1)
// 			if result != nil {
// 				return append([]*entities.Node{node}, result...)
// 			}
// 		}
// 	}
// 	return nil
// }

// func iterativeDeepeningSearch(root *entities.Node, target string) []*entities.Node {
// 	depth := 0
// 	for {
// 		fmt.Printf("Searching at depth %d\n", depth)
// 		result := dfs(root, target, depth)
// 		if result != nil {
// 			return result
// 		}
// 		depth++
// 	}
// }

// func IDS(ctx context.Context, siteURL, targetURL string, languageCode string) []string {
// 	depth := 0
// 	for {
// 		fmt.Printf("Searching at depth %d\n", depth)
// 		result := dfs(root, target, depth)
// 		if result != nil {
// 			return result
// 		}
// 		depth++
// 	}
// }

// // PageGraph represents a graph of Wikipedia pages that is built using
// // a bidirectional breadth-first search (forwards from a starting page and
// // backwards from an ending page).
// //
// // While searching, PageGraph runs two goroutines and, therefore, will at most
// // have two simultaneous API requests running against Wikipedia at a time.
// type PageGraph struct {
// 	// tree
// 	forwardTree entities.Node
// 	backwardTree entities.Node

// 	// map of page titles to their parent page title
// 	forward safeStringMap

// 	// queue of pages to search forwards from
// 	// forwardQueue []*entities.Node

// 	// map of page titles to their child page title
// 	backward safeStringMap

// 	// queue of pages to search backwards from
// 	// backwardQueue []*entities.Node
// }

// // NewPageGraph allocates and returns a PageGraph ready to search.
// func NewPageGraph(from, to string) PageGraph {
// 	return PageGraph{
// 		forwardTree : entities.Node{
// 			URL: from,
// 			Depth: 0,
// 			Visited: false,
// 		},
// 		backwardTree : entities.Node{
// 			URL: to,
// 			Depth: 0,
// 			Visited: false,
// 		},
// 		forward:       newSafeStringMap(),
// 		backward:      newSafeStringMap(),
// 	}
// }

// // func (pg *PageGraph) ClearPageGraph() {
// //     pg.forward = newSafeStringMap()
// //     pg.backward = newSafeStringMap()
// // }

// // Search takes starting and ending page titles and returns a short path of
// // links from the starting page to the ending page.
// func SearchIDS(from, to string, ctx context.Context, languageCode string) []string {
// 	pg := NewPageGraph(from, to)
// 	midpoint := make(chan string)

// 	go func() { midpoint <- pg.searchForwardIDS(from, ctx, languageCode) }()
// 	go func() { midpoint <- pg.searchBackwardIDS(to, ctx, languageCode) }()

// 	fmt.Println("midpoint: ", midpoint)

// 	return pg.path(<-midpoint)
// }

// func (pg *PageGraph) path(midpoint string) []string {
// 	path := []string{}

// 	// Build path from start to midpoint
// 	cursor := midpoint
// 	for len(cursor) > 0 {
// 		log.Printf("FOUND PATH FORWARD: %#v", cursor)
// 		path = append(path, cursor)
// 		cursor, _ = pg.forward.Get(cursor)
// 	}
// 	for i := 0; i < len(path)/2; i++ {
// 		swap := len(path) - i - 1
// 		path[i], path[swap] = path[swap], path[i]
// 	}

// 	PrintPath(path)

// 	// Pop off midpoint because the following loop adds it back in
// 	path = path[0 : len(path)-1]

// 	// Add path from midpoint to end
// 	cursor = midpoint
// 	for len(cursor) > 0 {
// 		log.Printf("FOUND PATH BACKWARDS: %#v", cursor)
// 		path = append(path, cursor)
// 		cursor, _ = pg.backward.Get(cursor)
// 	}

// 	PrintPath(path)

// 	return path
// }

// func (pg *PageGraph) searchForwardIDS(from string, ctx context.Context, languageCode string) string {
// 	pg.forward.Set(from, "")

// 	found := false
// 	foundLink := ""
// 	depth := 0
// 	for !found {
// 		foundLink, found = pg.searchForwardDLS(&pg.forwardTree, from, depth, ctx, languageCode)
// 		if (!found){
// 			depth++
// 		}
// 	}
// 	return foundLink
// }

// func (pg *PageGraph) searchForwardDLS(node *entities.Node, from string, depthLimit int, ctx context.Context, languageCode string) (foundLink string, found bool) {
// 	// basis tidak ditemukan
// 	if node == nil || depthLimit <= 0 {
// 		return "", false
// 	}

// 	// basis ditemukan
// 	if !node.Visited {
// 		node, _ =  scraping.GetWikiNodes(ctx, node.URL, languageCode)
// 		node.Visited = true
// 	}
// 	for _, child := range node.Children {
// 		if pg.checkForward(node.URL, child.URL){
// 			return child.URL, true
// 		}
// 	}

// 	// jika tidak menemukan atau tidak mencapai basis
// 	for _, child := range node.Children {
// 		result, found := pg.searchForwardDLS(child, from, depthLimit-1, ctx, languageCode)
// 		if found {
// 			return result, true
// 		}
// 	}
// 	return "", false
// }

// func (pg *PageGraph) checkForward(from, to string) (done bool) {
// 	_, exists := pg.forward.Get(to)
// 	if !exists {
// 		log.Printf("FORWARD %#v -> %#v", from, to)
// 		// "to" page doesn't have a path to the source yet.
// 		pg.forward.Set(to, from)
// 	}

// 	// If we now have a path to the destination, we're done!
// 	_, done = pg.backward.Get(to)
// 	return done
// }

// func (pg *PageGraph) searchBackwardIDS(to string, ctx context.Context, languageCode string) string {
// 	pg.backward.Set(to, "")

// 	found := false
// 	foundLink := ""
// 	depth := 0
// 	for !found {
// 		foundLink, found = pg.searchBackwardDLS(&pg.backwardTree, to, depth, ctx, languageCode)
// 		if (!found){
// 			depth++
// 		}
// 	}
// 	return foundLink
// }

// func (pg *PageGraph) searchBackwardDLS(node *entities.Node, to string, depthLimit int, ctx context.Context, languageCode string) (foundLink string, found bool) {
// 	// basis tidak ditemukan
// 	if node == nil || depthLimit <= 0 {
// 		return "", false
// 	}

// 	// basis ditemukan
// 	if !node.Visited {
// 		node, _ =  scraping.GetWikiNodes(ctx, node.URL, languageCode)
// 		node.Visited = true
// 	}
// 	for _, child := range node.Children {
// 		if pg.checkBackward(child.URL, node.URL){
// 			return child.URL, true
// 		}
// 	}

// 	// jika tidak menemukan atau tidak mencapai basis
// 	for _, child := range node.Children {
// 		result, found := pg.searchBackwardDLS(child, to, depthLimit-1, ctx, languageCode)
// 		if found {
// 			return result, true
// 		}
// 	}
// 	return "", false
// }

// func (pg *PageGraph) checkBackward(from, to string) (done bool) {
// 	_, exists := pg.backward.Get(from)
// 	if !exists {
// 		log.Printf("BACKWARD %#v -> %#v", from, to)
// 		// "from" page doesn't have a path to the destination yet.
// 		pg.backward.Set(from, to)
// 	}

// 	// If we now have a path to the source, we're done!
// 	_, done = pg.forward.Get(to)
// 	return done
// }

// // -- safeStringMap

// // safeStringMap is a helper type that wraps a map[string]string with
// // a sync.RWMutex.
// type safeStringMap struct {
// 	strings map[string]string
// 	sync.RWMutex
// }

// func newSafeStringMap() safeStringMap {
// 	return safeStringMap{map[string]string{}, sync.RWMutex{}}
// }

// func (m *safeStringMap) Get(key string) (value string, exists bool) {
// 	m.RLock()
// 	defer m.RUnlock()
// 	value, exists = m.strings[key]
// 	return
// }

// func (m *safeStringMap) Set(key, value string) {
// 	m.Lock()
// 	defer m.Unlock()
// 	m.strings[key] = value
// }

// func PrintPath(path []string) {
// 	result := strings.Join(path, " - ")
// 	fmt.Println(result)
// }
