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

func revLinkedList(now *ListNode) (*ListNode, *ListNode) {
	if now.Next == nil {
		return now, now
	}
	last, newHead := revLinkedList(now.Next)
	last.Next = now
	now.Next = nil
	return now, newHead
}
