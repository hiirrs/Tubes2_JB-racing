package entities

import (
	"fmt"
	"sync"
)

type Node struct {
	URL      string
	Parent   *Node
	Children []*Node
	Depth    int
	mutex    sync.Mutex
}

func (n *Node) AddChild(child *Node) {
	n.mutex.Lock()        // Lock before modifying
	defer n.mutex.Unlock()
	child.Parent = n
	child.Depth = n.Depth + 1
	n.Children = append(n.Children, child)
}

func (n *Node) AddChildIDS(child *Node) {
	for _, existingChild := range n.Children {
		if existingChild.URL == child.URL {
			return
		}
	}

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
