package main

import "fmt"

type Node struct {
	Value int
	Poped int
	Next  *Node
}
type Stack struct {
	Head *Node
}

func (s *Stack) Push(val int, poped int) {
	newNode := &Node{
		Value: val,
		Poped: poped,
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

func largestRectangleArea(heights []int) int {
	s := CreateNewStack()
	var res int = 0
	for i := 0; i < len(heights); i++ {
		cd := 1 // chieu dai
		for s.Head != nil && heights[i] <= s.Head.Value {
			cd += s.Head.Poped + 1
			res = max(res, (cd-1)*s.Head.Value)
			s.Pop()
		}
		/*
			fmt.Println(cd * heights[i])
			fmt.Println(res)
			fmt.Println("")
		*/
		res = max(res, cd*heights[i])
		s.Push(heights[i], cd-1)
	}
	if s.Head != nil {
		cd := 0 // chieu dai
		for s.Head != nil {
			/*
				fmt.Println(s.Head.Value)
				fmt.Println(s.Head.Poped)
			*/
			cd += s.Head.Poped + 1
			res = max(res, s.Head.Value*cd)
			s.Pop()
		}
	}
	return res
}

func main() {
	heights := []int{2, 1, 4, 5, 1, 3, 3}
	fmt.Println(largestRectangleArea(heights))
}
