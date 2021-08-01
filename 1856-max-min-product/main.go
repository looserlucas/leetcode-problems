package main

import (
	"fmt"
)

const MOD = 1e9 + 7

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

func maxSumMinProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	pre := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		pre[i+1] = pre[i] + nums[i]
	}

	s := CreateNewStack()
	for i := 0; i < len(nums); i++ {
		for s.Head != nil && nums[s.Head.Value] >= nums[i] {
			s.Pop()
		}
		if s.Head == nil {
			left[i] = 0
		} else {
			left[i] = s.Head.Value + 1
		}
		s.Push(i)
	}

	s = CreateNewStack()
	for i := len(nums) - 1; i >= 0; i-- {
		for s.Head != nil && nums[s.Head.Value] >= nums[i] {
			s.Pop()
		}
		if s.Head == nil {
			right[i] = len(nums) - 1
		} else {
			right[i] = s.Head.Value - 1
		}
		s.Push(i)
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		res = max(res, nums[i]*(pre[right[i]+1]-pre[left[i]]))
	}
	return res % MOD
}

func main() {
	nums := []int{1, 2, 3, 2}
	fmt.Println(maxSumMinProduct(nums))
}
