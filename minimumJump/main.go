package main

import "fmt"

func minimumJumps(forbidden []int, a int, b int, x int) int {
	dp := make([]int, 6001)
	for i := 0; i < len(forbidden); i++ {
		dp[forbidden[i]] = -1
	}
	var q [][]int
	q = append(q, []int{0, 0})
	i := 0
	for i < len(q) {
		s := q[i][0]
		if s == x {
			return dp[s]
		}
		if s > b {
			if dp[s-b] == 0 && q[i][1] != 1 {
				dp[s-b] = dp[s] + 1
				q = append(q, []int{s - b, 1})
			}
		}
		if s+a < 6001 {
			if dp[s+a] == 0 {
				dp[s+a] = dp[s] + 1
				q = append(q, []int{s + a, 0})
			}
		}
		i++
	}

	return -1
}

func main() {
	fmt.Println(minimumJumps([]int{8, 3, 16, 6, 12, 20}, 15, 13, 11))
	fmt.Println(minimumJumps([]int{14, 4, 18, 1, 15}, 3, 15, 9))
}
