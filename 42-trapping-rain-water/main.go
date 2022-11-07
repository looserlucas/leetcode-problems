package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}
type Stack struct {
	Head *Node
}

func (s *Stack) Push(val int) {
	newNode := &Node{
		Value: val,
	}

	if s.Head == nil {
		s.Head = newNode
	} else {
		tmp := s.Head
		s.Head = newNode
		newNode.Next = tmp
	}
}

func (s *Stack) Pop() int {
	if s.Head == nil {
		return -1
	}
	tmp := s.Head
	s.Head = tmp.Next
	return tmp.Value
}

func CreateNewStack() (s *Stack) {
	return &Stack{}
}

func trap(height []int) int {
	// preprocess
	if len(height) == 0 {
		return 0
	}
	var res int
	start := 0
	for i := range height {
		if height[i] != 0 {
			start = i
			break
		}
	}
	if start == 0 && height[start] == 0 {
		return 0
	}
	height = height[start:len(height)]

	//s := CreateNewStack()
	fmt.Println(height)
	return res
}

func main() {
	height := []int{0, 0, 0, 0, 0, 0, 0}
	trap(height)
}
