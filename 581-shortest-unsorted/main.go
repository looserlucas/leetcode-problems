package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}
type Stack struct {
	Head   *Node
	Length int
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
	s.Length++
}

func (s *Stack) Pop() int {
	if s.Head == nil {
		return -1
	}
	tmp := s.Head
	s.Head = tmp.Next
	s.Length--
	return tmp.Value
}

func CreateNewStack() (s *Stack) {
	return &Stack{
		Length: 0,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findUnsortedSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	s := CreateNewStack()
	l := len(nums)
	r := 0
	s.Push(0)
	for i := 1; i < len(nums); i++ {
		for s.Head != nil && nums[s.Head.Value] > nums[i] {
			l = min(l, s.Pop())
		}
		s.Push(i)
	}
	s1 := CreateNewStack()
	s1.Push(0)
	for i := len(nums) - 1; i >= 0; i-- {
		for s.Head != nil && nums[s.Head.Value] < nums[i] {
			r = max(r, s.Pop())
		}
		s.Push(i)
	}
	if r-l > 0 {
		return r - l + 1
	} else {
		return 0
	}
}

type Node struct {
	Value int
	Next  *Node
}
type Stack struct {
	Head   *Node
	Length int
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
	s.Length++
}

func (s *Stack) Pop() int {
	if s.Head == nil {
		return -1
	}
	tmp := s.Head
	s.Head = tmp.Next
	s.Length--
	return tmp.Value
}

func CreateNewStack() (s *Stack) {
	return &Stack{
		Length: 0,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findUnsortedSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	s := CreateNewStack()
	l := len(nums)
	r := 0
	s.Push(0)
	for i := 1; i < len(nums); i++ {
		for s.Head != nil && nums[s.Head.Value] > nums[i] {
			l = min(l, s.Pop())
		}
		s.Push(i)
	}
	s1 := CreateNewStack()
	s1.Push(0)
	for i := len(nums) - 1; i >= 0; i-- {
		for s.Head != nil && nums[s.Head.Value] < nums[i] {
			r = max(r, s.Pop())
		}
		s.Push(i)
	}
	if r-l > 0 {
		return r - l + 1
	} else {
		return 0
	}
}

func main() {
	nums := []int{1, 2, 5, 3, 4}
	fmt.Println(findUnsortedSubarray(nums))
}
