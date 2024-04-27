package getPath

import (
	"github.com/gocolly/colly/v2"
	"logic/internal/entities"
	scraping "logic/internal/tools"
)

func ScrapeToNodeIDS(root *entities.Node) {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if scraping.IsWikipediaArticle(link) {
			child := &entities.Node{URL: "https://en.wikipedia.org" + link}
			root.AddChildIDS(child)
		}
	})
	c.OnRequest(func(r *colly.Request) {})
	c.Visit(root.URL)
	c.Wait()
}

func IDS(root *entities.Node, target string, maxDepth int) []*entities.Node {
	for depth := 0; depth <= maxDepth; depth++ {
		result, found := DLS(root, target, depth)
		if found {
			return result
		}
	}
	return nil
}

func DLS(root *entities.Node, target string, depth int) ([]*entities.Node, bool) {
	visited := make(map[string]bool)
	return DLSR(root, target, depth, visited)
}

func DLSR(root *entities.Node, target string, depth int, visited map[string]bool) ([]*entities.Node, bool) {
	// fmt.Printf("Visiting node %s at depth %d\n", root.URL, root.Depth)
	// println(depth)
	scraping.ScrapeToNode(root, root.Depth)
	// printChildrenURLs(root)
	if root.URL == target {
		return []*entities.Node{root}, true
	}
	if depth <= 0 {
		return nil, false
	}
	if root.Depth >= depth {
		return nil, false
	}
	visited[root.URL] = true
	println(root.URL)
	println(len(root.Children))
	for _, link := range root.Children {
		if _, ok := visited[link.URL]; !ok {
			visited[link.URL] = true
			// println("link to inspect")
			// println(link.URL)
			result, found := DLSR(link, target, depth-1, visited)
			if found {
				return result, found // Return if the target is found
			}
		}
	}
	return nil, false
}

func Backtrack(nodes []*entities.Node, path []string) []string {
	node := nodes[0]
	if node.Parent == nil {
		println("terakhir")
		println(node.URL)
		path = append(path, node.URL)
		return path
	}
	println(node.URL)
	path = append(path, node.URL)
	println(path)
	nodes[0] = node.Parent
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
