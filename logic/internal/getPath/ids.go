package getPath

import (
	"context"
	"strings"
	"fmt"
	"log"
	"sync"

	"logic/internal/entities"
	"logic/internal/tools"
)

// PageGraph represents a graph of Wikipedia pages that is built using
// a bidirectional breadth-first search (forwards from a starting page and
// backwards from an ending page).
//
// While searching, PageGraph runs two goroutines and, therefore, will at most
// have two simultaneous API requests running against Wikipedia at a time.
type PageGraph struct {
	// tree
	forwardTree entities.Node
	backwardTree entities.Node

	// map of page titles to their parent page title
	forward safeStringMap

	// queue of pages to search forwards from
	// forwardQueue []*entities.Node

	// map of page titles to their child page title
	backward safeStringMap

	// queue of pages to search backwards from
	// backwardQueue []*entities.Node
}

// NewPageGraph allocates and returns a PageGraph ready to search.
func NewPageGraph(from, to string) PageGraph {
	return PageGraph{
		forwardTree : entities.Node{
			URL: from,
			Depth: 0,
			Visited: false,
		},
		backwardTree : entities.Node{
			URL: to,
			Depth: 0,
			Visited: false,
		},
		forward:       newSafeStringMap(),
		backward:      newSafeStringMap(),
	}
}

// func (pg *PageGraph) ClearPageGraph() {
//     pg.forward = newSafeStringMap()
//     pg.backward = newSafeStringMap()
// }

// Search takes starting and ending page titles and returns a short path of
// links from the starting page to the ending page.
func SearchIDS(from, to string, ctx context.Context, languageCode string) []string {
	pg := NewPageGraph(from, to)
	midpoint := make(chan string)

	go func() { midpoint <- pg.searchForwardIDS(from, ctx, languageCode) }()
	go func() { midpoint <- pg.searchBackwardIDS(to, ctx, languageCode) }()

	fmt.Println("midpoint: ", midpoint)

	return pg.path(<-midpoint)
}

func (pg *PageGraph) path(midpoint string) []string {
	path := []string{}

	// Build path from start to midpoint
	cursor := midpoint
	for len(cursor) > 0 {
		log.Printf("FOUND PATH FORWARD: %#v", cursor)
		path = append(path, cursor)
		cursor, _ = pg.forward.Get(cursor)
	}
	for i := 0; i < len(path)/2; i++ {
		swap := len(path) - i - 1
		path[i], path[swap] = path[swap], path[i]
	}

	PrintPath(path)

	// Pop off midpoint because the following loop adds it back in
	path = path[0 : len(path)-1]

	// Add path from midpoint to end
	cursor = midpoint
	for len(cursor) > 0 {
		log.Printf("FOUND PATH BACKWARDS: %#v", cursor)
		path = append(path, cursor)
		cursor, _ = pg.backward.Get(cursor)
	}

	PrintPath(path)

	return path
}

func (pg *PageGraph) searchForwardIDS(from string, ctx context.Context, languageCode string) string {
	pg.forward.Set(from, "")

	found := false
	foundLink := ""
	depth := 0
	for !found {
		foundLink, found = pg.searchForwardDLS(&pg.forwardTree, from, depth, ctx, languageCode)
		if (!found){
			depth++
		}
	}
	return foundLink
}

func (pg *PageGraph) searchForwardDLS(node *entities.Node, from string, depthLimit int, ctx context.Context, languageCode string) (foundLink string, found bool) {
	// basis tidak ditemukan
	if node == nil || depthLimit <= 0 {
		return "", false
	}

	// basis ditemukan
	if !node.Visited {
		node, _ =  scraping.GetWikiNodes(ctx, node.URL, languageCode)
		node.Visited = true
	}
	for _, child := range node.Children {
		if pg.checkForward(node.URL, child.URL){
			return child.URL, true
		}
	}

	// jika tidak menemukan atau tidak mencapai basis
	for _, child := range node.Children {
		result, found := pg.searchForwardDLS(child, from, depthLimit-1, ctx, languageCode)
		if found {
			return result, true
		}
	}
	return "", false
}

func (pg *PageGraph) checkForward(from, to string) (done bool) {
	_, exists := pg.forward.Get(to)
	if !exists {
		log.Printf("FORWARD %#v -> %#v", from, to)
		// "to" page doesn't have a path to the source yet.
		pg.forward.Set(to, from)
	}

	// If we now have a path to the destination, we're done!
	_, done = pg.backward.Get(to)
	return done
}

func (pg *PageGraph) searchBackwardIDS(to string, ctx context.Context, languageCode string) string {
	pg.backward.Set(to, "")

	found := false
	foundLink := ""
	depth := 0
	for !found {
		foundLink, found = pg.searchBackwardDLS(&pg.backwardTree, to, depth, ctx, languageCode)
		if (!found){
			depth++
		}
	}
	return foundLink
}

func (pg *PageGraph) searchBackwardDLS(node *entities.Node, to string, depthLimit int, ctx context.Context, languageCode string) (foundLink string, found bool) {
	// basis tidak ditemukan
	if node == nil || depthLimit <= 0 {
		return "", false
	}

	// basis ditemukan
	if !node.Visited {
		node, _ =  scraping.GetWikiNodes(ctx, node.URL, languageCode)
		node.Visited = true
	}
	for _, child := range node.Children {
		if pg.checkBackward(child.URL, node.URL){
			return child.URL, true
		}
	}

	// jika tidak menemukan atau tidak mencapai basis
	for _, child := range node.Children {
		result, found := pg.searchBackwardDLS(child, to, depthLimit-1, ctx, languageCode)
		if found {
			return result, true
		}
	}
	return "", false
}

func (pg *PageGraph) checkBackward(from, to string) (done bool) {
	_, exists := pg.backward.Get(from)
	if !exists {
		log.Printf("BACKWARD %#v -> %#v", from, to)
		// "from" page doesn't have a path to the destination yet.
		pg.backward.Set(from, to)
	}

	// If we now have a path to the source, we're done!
	_, done = pg.forward.Get(to)
	return done
}

// -- safeStringMap

// safeStringMap is a helper type that wraps a map[string]string with
// a sync.RWMutex.
type safeStringMap struct {
	strings map[string]string
	sync.RWMutex
}

func newSafeStringMap() safeStringMap {
	return safeStringMap{map[string]string{}, sync.RWMutex{}}
}

func (m *safeStringMap) Get(key string) (value string, exists bool) {
	m.RLock()
	defer m.RUnlock()
	value, exists = m.strings[key]
	return
}

func (m *safeStringMap) Set(key, value string) {
	m.Lock()
	defer m.Unlock()
	m.strings[key] = value
}

func PrintPath(path []string) {
	result := strings.Join(path, " - ")
	fmt.Println(result)
}