package binarytree

import (
	"fmt"
)

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

type BinaryTree struct {
	Root *Node
}

func (b *BinaryTree) insertRecursivily(newNode *Node) {
	var currentNode *Node

	if b.Root == nil {
		b.Root = currentNode
		fmt.Println("inserted root node")
		return
	}

	b.insert(newNode, b.Root)
}

func (b *BinaryTree) insert(newNode *Node, currentNode *Node) {

	if newNode.Value > currentNode.Value {
		if currentNode.Right == nil {
			fmt.Printf("inserted the new node in the right of %v\n", currentNode.Value)
			currentNode.Right = newNode
			newNode.Parent = currentNode
			return
		}

		fmt.Printf("called b.insert in the right node of %c\n", currentNode.Value)
		b.insert(newNode, currentNode.Right)
	}

	if newNode.Value < currentNode.Value {
		// calls insert wiht the currentNode.left
		if currentNode.Left == nil {
			fmt.Printf("inserted the new node in the left of %v\n", currentNode.Value)
			currentNode.Left = newNode
			newNode.Parent = currentNode
			return
		}
		fmt.Printf("called b.insert in the left node of %c\n", currentNode.Value)
		b.insert(newNode, currentNode.Left)
	}
}

func main() {
	var myTree *BinaryTree

	for i := 0; i < 5; i++ {
		var node Node
		node.Value = 10
		myTree.insertRecursivily(&node)
	}
}
