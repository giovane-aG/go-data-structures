package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Down  *Node
	Value int
}

type Stack struct {
	Size int
	Top  *Node
}

func (s *Stack) push(node *Node) {
	s.Size++
	node.Down = s.Top
	s.Top = node
}

func (s *Stack) pop() *Node {

	if s.Size == 0 {
		fmt.Println("stack is already empty")
		return nil
	}

	s.Size--
	popped := s.Top
	s.Top = s.Top.Down
	return popped
}

func (s *Stack) peak() {
	fmt.Printf("Top node -> %v\n", s.Top.Value)
}

func (s Stack) printStack() {
	node := s.Top
	println("-----Stack-----")
	for node.Down != nil {
		fmt.Println(node.Value)
		node = node.Down
	}

	fmt.Println(node.Value)
}

func createStack() Stack {
	var s Stack
	s.Size = 0
	return s
}

func main() {
	stack := createStack()
	fmt.Println(stack)
	stack.pop()

	fmt.Println("Inserting values...")
	for i := 0; i < 10; i++ {
		var node Node
		node.Value = i + rand.Int()*rand.Int()
		stack.push(&node)
	}

	stack.peak()
	stack.printStack()

	for i := 0; i < 5; i++ {
		_ = stack.pop()
	}

	stack.peak()
	stack.printStack()
}
