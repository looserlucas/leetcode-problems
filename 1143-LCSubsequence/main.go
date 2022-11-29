package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func longestCommonSubsequence(text1 string, text2 string) int {
	var dp [][]int
	for i := 0; i <= len(text1); i++ {
		var tmp = make([]int, len(text2)+1)
		dp = append(dp, tmp)
	}

	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func main() {
	//fmt.Println(longestCommonSubsequence("abcde", "ace"))
	//fmt.Println(longestCommonSubsequence("abc", "abc"))
	//fmt.Println(longestCommonSubsequence("abc", "def"))
	fmt.Println(longestCommonSubsequence("ezupkr", "ubmrapg"))
}
