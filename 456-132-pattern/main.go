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
func find132pattern(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	s := CreateNewStack()
	s3 := -9999999999
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < s3 {
			return true
		}
		for s.Head != nil && nums[i] > s.Head.Value {
			s3 = s.Pop()
		}
		s.Push(nums[i])
	}
	return false
}
func main() {
	nums := []int{1, 0, 1, -4, -3}
	fmt.Println(find132pattern(nums))
}
