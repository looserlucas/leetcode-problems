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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxWidthRamp(nums []int) int {
	s := CreateNewStack()
	for i := 0; i < len(nums); i++ {
		if s.Head == nil || nums[s.Head.Value] > nums[i] {
			s.Push(i)
		}
	}

	res := 0
	for j := len(nums) - 1; j >= 0; j-- {
		for s.Head != nil && nums[j]-nums[s.Head.Value] >= 0 {
			res = max(res, j-s.Pop())
		}
	}
	return res
}

func main() {
	nums := []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}
	fmt.Println(maxWidthRamp(nums))
}
