package main

import (
	"fmt"
)

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

// Node functions

func (n Node) hasLeftNode() bool {
	return n.Left != nil
}

func (n Node) hasRightNode() bool {
	return n.Right != nil
}

func (n Node) isLeftSon() bool {
	if n.Parent.Left == &n {
		return true
	}
	return false
}

func (n Node) isRightSon() bool {
	if n.Parent.Right == &n {
		return true
	}
	return false
}

func (n Node) successor() *Node {
	node := successorRecursively(n.Right)
	return node
}

func successorRecursively(node *Node) *Node {
	if node.hasLeftNode() {
		successorRecursively(node.Left)
	}
	return node
}

// BTree functions

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
			// fmt.Printf("inserted %v in the right of %v\n", newNode.Value, currentNode.Value)
			currentNode.Right = newNode
			newNode.Parent = currentNode
			return
		}

		// fmt.Printf("called b.insert in the right node of %v\n", currentNode.Value)
		b.insert(newNode, currentNode.Right)
	}

	if newNode.Value < currentNode.Value {
		// calls insert wiht the currentNode.left
		if currentNode.Left == nil {
			// fmt.Printf("inserted %v in the left of %v\n", newNode.Value, currentNode.Value)
			currentNode.Left = newNode
			newNode.Parent = currentNode
			return
		}
		// fmt.Printf("called b.insert in the left node of %v\n", currentNode.Value)
		b.insert(newNode, currentNode.Left)
	}
}

func (b *BinaryTree) remove(value int) {
	var currentNode *Node = b.Root
	b.removeRecursively(value, currentNode)
}

func (b *BinaryTree) removeRecursively(value int, currentNode *Node) {
	if value > currentNode.Value {
		b.removeRecursively(value, currentNode.Right)
	}
	if value < currentNode.Value {
		b.removeRecursively(value, currentNode.Left)
	}

	if value == currentNode.Value {
		if currentNode == b.Root {
			b.Root = nil

		}

		if !currentNode.hasRightNode() && !currentNode.hasLeftNode() {
			currentNode = nil
		}

		if currentNode.hasRightNode() && !currentNode.hasLeftNode() {
			currentNode.Right.Parent = currentNode.Parent

			if currentNode.isLeftSon() {
				currentNode.Parent.Left = currentNode.Right
			} else {
				currentNode.Parent.Right = currentNode.Right
			}
		}

		if currentNode.hasLeftNode() && !currentNode.hasRightNode() {
			currentNode.Left.Parent = currentNode.Parent

			if currentNode.isRightSon() {
				currentNode.Parent.Right = currentNode.Left
			} else {
				currentNode.Parent.Left = currentNode.Left
			}
		}

		// TODO: node with two sons

	}
}

func (b *BinaryTree) inOrderTransversal() {
	node := b.Root
	b.inOrderTransversalRecursively(node)
}

func (b *BinaryTree) inOrderTransversalRecursively(node *Node) {
	if node != nil {
		b.inOrderTransversalRecursively(node.Left)
		fmt.Println(node.Value)
		b.inOrderTransversalRecursively(node.Right)
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

	var node1 Node
	var node2 Node
	var node3 Node
	var node4 Node
	var node5 Node
	var node6 Node
	var node7 Node
	var node8 Node

	node1.Value = 10
	node2.Value = 8
	node3.Value = 9
	node4.Value = 12
	node5.Value = 11
	node6.Value = 16
	node7.Value = 14
	node8.Value = 19

	tree.insertRecursively(&node1)
	tree.insertRecursively(&node2)
	tree.insertRecursively(&node3)
	tree.insertRecursively(&node4)
	tree.insertRecursively(&node5)
	tree.insertRecursively(&node6)
	tree.insertRecursively(&node7)
	tree.insertRecursively(&node8)

	// tree.inOrderTransversal()

	tree.remove(16)
	tree.inOrderTransversal()
}
