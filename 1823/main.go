package main

import "fmt"

type node struct {
	Val  int
	Next *node
	Prev *node
}

func findTheWinner(n int, k int) int {
	head := &node{
		Val: 1,
	}
	tail := head
	len := n
	for i := 2; i <= n; i++ {
		tmp := &node{
			Val:  i,
			Prev: tail,
		}
		tail.Next = tmp
		tail = tmp
	}
	tail.Next = head
	head.Prev = tail
	now := head
	var res int = head.Val
	for len > 1 {
		//fmt.Print(now.Val, " /")
		count := k - 1
		for count > 0 {
			now = now.Next
			count--
		}
		//fmt.Println(now.Prev, now.Val, now.Next.Val)
		// remove now
		next := now.Next
		prev := now.Prev

		prev.Next = next
		next.Prev = prev
		now.Prev = nil
		now.Next = nil

		now = next
		res = now.Val
		len--
	}

	return res
}

func main() {
	fmt.Println(findTheWinner(5, 2))
	fmt.Println(findTheWinner(6, 5))
}
