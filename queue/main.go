package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value    float32 `json:"value"`
	Next     *Node   `json:"next"`
	Previous *Node   `json:"previous"`
}

func (node Node) printNode() {
	fmt.Printf("----Node %v----\n", node.Value)
	fmt.Printf("next: %v\n", node.Next.Value)

	if node.Previous != nil {
		fmt.Printf("previous: %v\n", node.Previous.Value)
	}
}

type Queue struct {
	size  int
	first *Node
	last  *Node
}

func (queue Queue) isEmpty() bool {
	return queue.size == 0
}

func (queue *Queue) queue(node *Node) {

	if queue.isEmpty() {
		queue.first = node
		queue.size++
		return
	}

	actualNode := queue.first

	for actualNode.Next != nil {
		actualNode = actualNode.Next
	}

	actualNode.Next = node
	node.Previous = actualNode
	queue.last = node
	queue.size++
}

func (queue *Queue) unqueue() (*Node, error) {
	var unqueuedNode *Node

	if queue.isEmpty() {
		return nil, errors.New("queue is already empty")
	}

	unqueuedNode = queue.first

	if queue.size > 1 {
		queue.first = queue.first.Next
		queue.first.Previous = nil
	} else {
		queue.first = nil
	}

	queue.size--

	return unqueuedNode, nil
}

func NewQueue() *Queue {
	var newQueue Queue
	newQueue.size = 0
	return &newQueue
}

func main() {
	q := NewQueue()
	// fmt.Println(q)

	for i := 0; i < 10; i++ {
		var node Node
		node.Value = float32(i)
		q.queue(&node)
	}

	node := q.first

	for node.Next != nil {
		node.printNode()
		node = node.Next
	}

	q.unqueue()
}
