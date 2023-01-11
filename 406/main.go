package main

import (
	"fmt"
	"sort"
)

type node struct {
	Val  []int
	Next *node
	Prev *node
}

type queue struct {
	Head *node
	Tail *node
}

func (q *queue) addToQueueAtK(n []int, k int) {
	tmp := &node{
		Val: n,
	}
	if q.Head == nil {
		q.Head = tmp
		q.Tail = tmp
		return
	}

	now := q.Head
	count := 0
	for count < k {
		now = now.Next
		count++
	}
	if now == nil {
		q.Tail.Next = tmp
		tmp.Prev = q.Tail
		q.Tail = tmp
	} else if now == q.Head {
		tmp.Next = q.Head
		q.Head.Prev = tmp
		q.Head = tmp
	} else {
		prev := now.Prev
		prev.Next = tmp
		tmp.Next = now
		tmp.Prev = prev
		now.Prev = tmp
	}
}

func reconstructQueue(a [][]int) [][]int {
	sort.SliceStable(a, func(i, j int) bool {
		if a[i][0] == a[j][0] {
			return a[i][1] < a[j][1]
		}
		return a[i][0] > a[j][0]
	})
	var res [][]int
	var q *queue = &queue{
		Head: nil,
		Tail: nil,
	}
	for i := 0; i < len(a); i++ {
		q.addToQueueAtK(a[i], a[i][1])
	}
	now := q.Head
	for i := 0; i < len(a); i++ {
		res = append(res, now.Val)
		now = now.Next
	}
	return res
}

func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
	fmt.Println(reconstructQueue([][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}))
}
