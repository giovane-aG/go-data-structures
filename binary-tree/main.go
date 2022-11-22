package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

func (n Node) hasLeftNode() bool {
	return n.Left != nil
}

func (n Node) hasRightNode() bool {
	return n.Right != nil
}

type BinaryTree struct {
	Root *Node
}

func (b *BinaryTree) insertRecursively(newNode *Node) {
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
			fmt.Printf("inserted %v in the right of %v\n", newNode.Value, currentNode.Value)
			currentNode.Right = newNode
			newNode.Parent = currentNode
			return
		}

		fmt.Printf("called b.insert in the right node of %v\n", currentNode.Value)
		b.insert(newNode, currentNode.Right)
	}

	if newNode.Value < currentNode.Value {
		// calls insert wiht the currentNode.left
		if currentNode.Left == nil {
			fmt.Printf("inserted %v in the left of %v\n", newNode.Value, currentNode.Value)
			currentNode.Left = newNode
			newNode.Parent = currentNode
			return
		}
		fmt.Printf("called b.insert in the left node of %v\n", currentNode.Value)
		b.insert(newNode, currentNode.Left)
	}
}

func (b *BinaryTree) inOrderTransversal() {
	node := b.Root
	b.inOrderTransversalRecursivily(node)
}

func (b *BinaryTree) inOrderTransversalRecursively(node *Node) {
	if node != nil {
		b.inOrderTransversalRecursivily(node.Left)
		fmt.Println(node.Value)
		b.inOrderTransversalRecursivily(node.Right)
	}
}

func CreateTree(rootValue int) *BinaryTree {
	var tree BinaryTree
	var root Node
	root.Value = rootValue
	tree.Root = &root
	return &tree
}

func main() {
	tree := CreateTree(10)

	for i := 0; i < 10; i++ {
		var node Node
		node.Value = rand.Intn(50)
		tree.insertRecursivily(&node)
	}

	tree.inOrderTransversal()
}
