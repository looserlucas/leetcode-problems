package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// the idea is to turn our first linkedlist into 2 odd - even linked list, then remerge them
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	headOdd := head
	nowOdd := headOdd
	lastOdd := headOdd
	headEven := head.Next
	nowEven := headEven
	for nowOdd != nil && nowEven != nil {
		if nowOdd.Next != nil {
			nowOdd.Next = nowOdd.Next.Next
			lastOdd = nowOdd
			nowOdd = nowOdd.Next
		}
		if nowEven.Next != nil {
			nowEven.Next = nowEven.Next.Next
			nowEven = nowEven.Next
		}
	}
	if nowEven != nil {
		nowEven.Next = nil
	}
	if nowOdd == nil {
		lastOdd.Next = headEven
	} else {
		nowOdd.Next = headEven
	}
	return head
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

func main() {
	/*
		printList(oddEvenList(arrToLinkedList([]int{1, 2, 3, 4, 5})))
		printList(oddEvenList(arrToLinkedList([]int{2, 1, 3, 5, 6, 4, 7})))
	*/
	printList(oddEvenList(arrToLinkedList([]int{1, 2})))
}
