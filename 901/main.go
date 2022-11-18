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

type StockSpanner struct {
	S *Stack
	P []int
	C []int
}

func Constructor() StockSpanner {
	s := CreateNewStack()
	var c, p []int
	return StockSpanner{
		S: s,
		C: c,
		P: p,
	}
}

func (this *StockSpanner) Next(price int) int {
	var res = 1
	for this.S.Head != nil && price >= this.P[this.S.Head.Value] {
		res += this.C[this.S.Pop()]
	}
	this.C = append(this.C, res)
	this.P = append(this.P, price)
	this.S.Push(len(this.P) - 1)
	return res
}

func main() {
	st := Constructor()
	next := []int{100, 80, 60, 70, 60, 75, 85}
	for _, v := range next {
		fmt.Println(st.Next(v))
	}
}
