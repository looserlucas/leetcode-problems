package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	helper(dummy, 0, left, right)
	return dummy.Next
}

func helper(now *ListNode, count int, left, right int) (newHead *ListNode, lastNodeNext *ListNode, lastNode *ListNode) {
	if right == count {
		next := now.Next
		now.Next = nil
		return now, next, now
	} else if left <= count && count < right {
		newHead, lastNodeNext, lastNode := helper(now.Next, count+1, left, right)
		lastNode.Next = now
		now.Next = nil
		return newHead, lastNodeNext, now
	} else if count == left-1 {
		newHead, lastNodeNext, lastNode := helper(now.Next, count+1, left, right)
		now.Next = newHead
		lastNode.Next = lastNodeNext
		return nil, nil, nil
	} else {
		helper(now.Next, count+1, left, right)
		return nil, nil, nil
	}
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
	head := arrToLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	printList(reverseBetween(head, 1, 1))
}
