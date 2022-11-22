package main

import "fmt"

func minCost(colors string, neededTime []int) int {
	var res int
	var now = 0
	for i := 1; i < len(colors); i++ {
		if colors[i] == colors[now] {
			if neededTime[i] > neededTime[now] {
				res += neededTime[now]
				now = i
			} else {
				res += neededTime[i]
			}
		} else {
			now = i
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
