package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	save := 0
	var res *ListNode
	if t := l1.Val + l2.Val; t < 10 {
		res = &ListNode{
			Val: t,
		}
	} else {
		res = &ListNode{
			Val: t - 10,
		}
		save = 1
	}
	n1 := l1.Next
	n2 := l2.Next
	now := res
	for n1 != nil || n2 != nil {
		var num1, num2 int
		if n1 == nil {
			num1 = 0
		} else {
			num1 = n1.Val
		}
		if n2 == nil {
			num2 = 0
		} else {
			num2 = n2.Val
		}
		if t := num1 + num2 + save; t < 10 {
			tmp := &ListNode{
				Val: t,
			}
			now.Next = tmp
			now = tmp
			save = 0
		} else {
			tmp := &ListNode{
				Val: t - 10,
			}
			save = 1
			now.Next = tmp
			now = tmp
		}
		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}
	if save > 0 {
		tmp := &ListNode{
			Val: 1,
		}
		save = 1
		now.Next = tmp
		now = tmp
	}

	return res
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
func main() {
	addTwoNumbers(arrToLinkedList([]int{9, 9, 9, 9, 9, 9, 9}), arrToLinkedList([]int{9, 9, 9, 9}))
	addTwoNumbers(arrToLinkedList([]int{0}), arrToLinkedList([]int{0}))
	addTwoNumbers(arrToLinkedList([]int{2, 4, 3}), arrToLinkedList([]int{5, 6, 4}))
	/*
		list := arrToLinkedList([]int{1, 2})
		a, _ := revLinkedList(list, list.Next)
		printList(a)
	*/
}
