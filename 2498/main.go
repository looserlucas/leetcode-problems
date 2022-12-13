package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxJump(stones []int) int {
	res := stones[1] - stones[0]
	for i := 3; i < len(stones); i++ {
		res = max(res, stones[i]-stones[i-2])
	}
	for i := 2; i < len(stones); i++ {
		res = max(res, stones[i]-stones[i-2])
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
