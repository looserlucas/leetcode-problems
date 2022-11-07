package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func arrToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	res := &ListNode{
		Val: arr[0],
	}
	now := res
	for i := 1; i < len(arr); i++ {
		tmp := &ListNode{
			Val: arr[i],
		}
		now.Next = tmp
		now = tmp
	}
	return res
}
func printList(node *ListNode) {
	now := node
	for now != nil {
		fmt.Print(now.Val, " ")
		now = now.Next
	}
	fmt.Println()
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
