package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func bestTeamScore(scores []int, ages []int) int {
	var dp = make([]int, len(scores))
	var sa [][]int
	for i := 0; i < len(scores); i++ {
		sa = append(sa, []int{ages[i], scores[i]})
	}
	sort.SliceStable(sa, func(i, j int) bool {
		if sa[i][0] > sa[j][0] {
			return true
		} else if sa[i][0] == sa[j][0] {
			return sa[i][1] > sa[j][1]
		}
		return false
	})

	var res = -1
	for i := 0; i < len(sa); i++ {
		dp[i] = sa[i][1]
		for j := 0; j < i; j++ {
			if sa[j][1] >= sa[i][1] {
				dp[i] = max(dp[i], dp[j]+sa[i][1])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
