package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	h, _ := revLinkedList(head, head.Next)
	return h
}

func revLinkedList(prev *ListNode, now *ListNode) (head *ListNode, end *ListNode) {
	if now.Next == nil {
		now.Next = prev
		prev.Next = nil
		return now, prev
	} else {
		head, node := revLinkedList(now, now.Next)
		node.Next = now
		now.Next = prev
		prev.Next = nil
		return head, prev
	}
}

func main() {
	fmt.Println("vim-go")
}
