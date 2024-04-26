package entities

import "fmt"

type Node struct {
	URL     string
	Parent  *Node
	Children []*Node
	Depth int
}

func (n *Node) AddChild(child *Node) {
	child.Parent = n
	child.Depth = n.Depth + 1
	n.Children = append(n.Children, child)
}

func PrintTree(node *Node, level int) {
	if node == nil {
		return
	}

	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}

	fmt.Println(node.URL)

	for _, child := range node.Children {
		PrintTree(child, level+1)
	}
}
