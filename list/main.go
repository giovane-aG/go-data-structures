package main

import "fmt"

type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

type List struct {
	Size  uint
	First *Node
	Last  *Node
}

func (l *List) insertAtTheEnd(newNode *Node) {

	if l.Size == 0 {
		l.First = newNode
		l.Last = newNode
		l.Size = l.Size + 1
		return
	}

	l.Last.Next = newNode
	newNode.Previous = l.Last
	l.Last = newNode
	l.Size = l.Size + 1
}

func (l *List) remove(index uint) {
	/*
		verify if the given index is closer to the end or the begining
		if closer to the end, iterate backwards in the list
		else iterate foward from the begining
	*/

	var actualNode *Node

	// index == first position
	if index == 0 {
		actualNode = l.First
		l.First = actualNode.Next
		l.First.Previous = nil
		actualNode.Next = nil
		l.Size--
		return
	}

	// index == last position
	if index == uint(l.Size)-1 {
		actualNode = l.Last
		l.Last = actualNode.Previous
		l.Last.Next = nil
		actualNode.Previous = nil
		l.Size--
		return
	}

	var limit uint
	var differenceBeginingIndex uint = index - 0
	var differenceEndIndex uint = l.Size - 1 - index

	if differenceBeginingIndex > differenceEndIndex {

		actualNode = l.Last
		limit = differenceEndIndex

		var i uint
		for i = 0; i < limit; i++ {
			actualNode = actualNode.Previous
		}
	} else {

		actualNode = l.First
		limit = differenceBeginingIndex
		var i uint
		for i = 0; i < limit; i++ {
			actualNode = actualNode.Next
		}
	}

	actualNode.Previous.Next = actualNode.Next
	actualNode.Next.Previous = actualNode.Previous
	actualNode.Previous = nil
	actualNode.Next = nil
	l.Size--
}

func CreateList() List {
	var list List
	list.Size = 0
	list.First = nil
	list.Last = nil

	return list
}

func main() {
	var newList List = CreateList()

	newNode := createNode(6)
	newNode2 := createNode(15)
	newNode3 := createNode(17)
	newNode4 := createNode(21)
	newNode5 := createNode(0)
	newNode6 := createNode(-12334)

	newList.insertAtTheEnd(newNode)
	newList.insertAtTheEnd(newNode2)
	newList.insertAtTheEnd(newNode3)
	newList.insertAtTheEnd(newNode4)
	newList.insertAtTheEnd(newNode5)
	newList.insertAtTheEnd(newNode6)

	printList(newList)
	newList.remove(5)
	newList.remove(0)

	printList(newList)
}

func printList(l List) {
	fmt.Println("-----List-----")
	fmt.Printf("First -> %v\n", l.First.Value)
	fmt.Printf("Last -> %v\n", l.Last.Value)
	fmt.Printf("Size -> %v\n", l.Size)

	print("[")
	var node *Node = l.First
	for i := 0; i <= int(l.Size)-1; i++ {

		if i != int(l.Size)-1 {
			fmt.Printf("%v, ", node.Value)
		} else {
			fmt.Printf("%v", node.Value)
		}

		node = node.Next
	}
	print("]\n")
}

func createNode(value int) *Node {
	var newNode Node
	newNode.Value = value
	return &newNode
}
