package newGetPath

import (
	"context"

	"logic/internal/entities"
	scraping "logic/internal/tools"
)

// PageGraph represents a graph of Wikipedia pages that is built using
// a unidirectional breadth-first search from a starting page.
type PageGraph struct {
	// tree
	root entities.Node

	// map of page titles to their parent page title
	parents safeStringMap

	// queue of pages to search from
	queue []*entities.Node
}

// NewPageGraph allocates and returns a PageGraph ready to search.
func NewPageGraph(from string) PageGraph {
	return PageGraph{
		root: entities.Node{
			URL:     from,
			Depth:   0,
			Visited: false,
		},
		parents: newSafeStringMap(),
		queue:   make([]*entities.Node, 0),
	}
}

// Search takes a starting page title and returns a path of
// links from the starting page to the ending page.
func Search(from, to string, ctx context.Context, languageCode string) []string {
	pg := NewPageGraph(from)
	return pg.searchIDS(from, to, ctx, languageCode)
}

func (pg *PageGraph) searchIDS(from, to string, ctx context.Context, languageCode string) []string {
	pg.parents.Set(from, "")
	pg.queue = append(pg.queue, &pg.root)

	var currentNode *entities.Node
	for len(pg.queue) > 0 {
		currentNode = pg.queue[0]
		pg.queue = pg.queue[1:]

		if currentNode.URL == to {
			break
		}

		if !currentNode.Visited {
			currentNode, _ = scraping.GetWikiNodes(ctx, currentNode.URL, languageCode)
			currentNode.Visited = true
		}

		for _, child := range currentNode.Children {
			if _, exists := pg.parents.Get(child.URL); !exists {
				pg.parents.Set(child.URL, currentNode.URL)
				pg.queue = append(pg.queue, child)
			}
		}
	}

	return pg.buildPath(to)
}

func (pg *PageGraph) buildPath(to string) []string {
	path := []string{}
	cursor := to
	for len(cursor) > 0 {
		path = append(path, cursor)
		var exists bool
		cursor, exists = pg.parents.Get(cursor)
		if !exists {
			break
		}
	}

	// Reverse path to start from the source
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}

	PrintPath(path)
	return path
}

// -- safeStringMap and other types/functions remain unchanged from the original code
