package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Value byte
	Next  *Node
}
type Stack struct {
	Head *Node
}

func (s *Stack) Push(val byte) {
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

func (s *Stack) Pop() byte {
	if s.Head == nil {
		return byte(0)
	}
	tmp := s.Head
	s.Head = tmp.Next
	return tmp.Value
}

func CreateNewStack() (s *Stack) {
	return &Stack{}
}

func isUpperCase(c byte) bool {
	return int(c) >= int('A') && int(c) <= int('Z')
}
func isLowerCase(c byte) bool {
	return int(c) >= int('a') && int(c) <= int('z')
}

func makeGood(s string) string {
	if len(s) == 1 {
		return s
	}
	st := CreateNewStack()
	for i := 0; i < len(s); i++ {
		if st.Head == nil {
			st.Push(s[i])
			continue
		}
		if isUpperCase(s[i]) && isLowerCase(st.Head.Value) && strings.EqualFold(strings.ToLower(string(s[i])), strings.ToLower(string(st.Head.Value))) {
			st.Pop()
			continue
		}
		if isLowerCase(s[i]) && isUpperCase(st.Head.Value) && strings.EqualFold(strings.ToLower(string(s[i])), strings.ToLower(string(st.Head.Value))) {
			st.Pop()
			continue
		}
		st.Push(s[i])
	}
	var res string = ""
	for st.Head != nil {
		res = string(st.Pop()) + res
	}
	return res
}

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood("abBAcC"))
	fmt.Println(makeGood("s"))
	fmt.Println(makeGood("mC"))
}
