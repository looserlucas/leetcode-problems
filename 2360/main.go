package main

import "fmt"

type stack []int

func (s *stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}

	x := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return x
}

func (s *stack) Push(x int) {
	(*s) = append((*s), x)
}

var id int
var res int

var s stack
var ids []int
var low []int
var inStack []bool

func longestCycle(edges []int) int {
	ids = nil
	low = nil
	inStack = nil
	res = 0
	id = 0
	s = nil
	for i := 0; i < len(edges); i++ {
		ids = append(ids, -1)
		low = append(low, 0)
		inStack = append(inStack, false)
	}

	for i := 0; i < len(edges); i++ {
		if ids[i] == -1 {
			dfs(edges, i)
		}
	}
	if res == 1 {
		return -1
	} else {
		return res
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func dfs(edges []int, i int) {
	//bookeeping
	s.Push(i)
	inStack[i] = true
	ids[i] = id
	low[i] = id
	id++

	//min low
	v := edges[i]
	if v != -1 {
		if ids[v] == -1 {
			dfs(edges, v)
		}
		if inStack[v] {
			low[i] = min(low[i], low[v])
		}
	}

	//check for SCC
	if low[i] == ids[i] {
		var count int
		for len(s) > 0 {
			x := s.Pop()
			inStack[x] = false
			count++
			if x == i {
				break
			}
		}
		if count > res {
			res = count
		}
	}
}

func main() {
	fmt.Println(longestCycle([]int{3, 3, 4, 2, 3}))
	fmt.Println(longestCycle([]int{2, -1, 3, 1}))
}
