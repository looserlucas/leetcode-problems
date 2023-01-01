package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var head, tail *ListNode

func mergeNodes(h *ListNode) *ListNode {
	head, tail = nil, nil
	merge(h, false)
	return head
}

func merge(now *ListNode, new bool) {
	if now == nil {
		return
	}
	if now.Val == 0 {
		if head == nil {
			head = now.Next
			tail = now.Next
			merge(now.Next, true)
		} else {
			tail.Next = now.Next
			tail = now.Next
			merge(now.Next, true)
		}
	} else {
		if !new {
			tail.Val += now.Val
		}
		merge(now.Next, false)
	}
}
func main() {
	fmt.Println("vim-go")
}
