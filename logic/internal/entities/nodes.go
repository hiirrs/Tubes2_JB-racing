package entities

import "fmt"

// Node represents a node in the tree structure
type Node struct {
	URL     string
	Children []*Node
	Depth int
}

// AddChild adds a child node to the current node
func (n *Node) AddChild(child *Node) {
	child.Depth = n.Depth + 1
	n.Children = append(n.Children, child)
	// fmt.Println("child created")
}

// PrintTree recursively prints the tree structure
func PrintTree(node *Node, level int) {
	if node == nil {
		return
	}

	// Indent based on the level
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}

	// Print the URL of the current node
	fmt.Println(node.URL)

	// Recursively print child nodes
	for _, child := range node.Children {
		PrintTree(child, level+1)
	}
}
