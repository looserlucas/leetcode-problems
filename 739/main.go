package main

import "fmt"

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

func dailyTemperatures(t []int) []int {
	s := CreateNewStack()
	var res = make([]int, len(t))
	for i := 0; i < len(t); i++ {
		for s.Head != nil && t[s.Head.Value] < t[i] {
			now := s.Pop()
			res[now] = i - now
		}
		s.Push(i)
	}
	return res
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}
