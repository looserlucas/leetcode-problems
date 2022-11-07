package main

import (
	"fmt"
	"math"
)

var sessionTime int
var tasks []int
var dp [][]int

func minSessions(t []int, s int) int {
	sessionTime = s
	tasks = t
	var target int
	for i := 0; i < len(t); i++ {
		var mask = 1 << i
		target = target | mask
	}
	dp = nil
	for i := 0; i <= target; i++ {
		var temp []int
		for j := 0; j <= sessionTime; j++ {
			temp = append(temp, -1)
		}
		dp = append(dp, temp)
	}
	return dfs(target, 0, sessionTime) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func dfs(target int, done int, left int) int {
	if dp[done][left] != -1 {
		return dp[done][left]
	}
	if done == target {
		return 0
	}

	res := math.MaxInt
	for i := 0; i < len(tasks); i++ {
		var mask = 1 << i
		if mask&done == 0 {
			newDone := mask | done
			if tasks[i] > left {
				res = min(res, dfs(target, newDone, sessionTime-tasks[i])+1)
			} else {
				res = min(res, dfs(target, newDone, left-tasks[i]))
			}
		}
	}
	dp[done][left] = res
	return res
}

func main() {
	fmt.Println(minSessions([]int{1, 2, 3}, 3))
	fmt.Println(minSessions([]int{3, 1, 3, 1, 1}, 8))
	fmt.Println(minSessions([]int{1, 2, 3, 4, 5}, 15))
	fmt.Println(minSessions([]int{1, 2, 3, 4, 5, 6}, 15))
	fmt.Println(minSessions([]int{3, 2, 3, 7, 5, 2, 2, 10, 9, 1, 10}, 11))
}
