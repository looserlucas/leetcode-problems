package main

import "fmt"

const MAX_INT = int(1e9)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func minDistance(word1 string, word2 string) int {
	word1 = " " + word1
	word2 = " " + word2

	var dp [][]int
	for i := 0; i < len(word1); i++ {
		var tmp []int
		for j := 0; j < len(word2); j++ {
			tmp = append(tmp, MAX_INT)
		}
		dp = append(dp, tmp)
	}

	for i := 0; i < len(word1); i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(word2); i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				//                  insert          delete       replace
				dp[i][j] = 1 + min(dp[i][j-1], min(dp[i-1][j], dp[i-1][j-1]))
			}
		}
	}
	return dp[len(word1)-1][len(word2)-1]
}
func main() {
	word1 := "intention"
	word2 := "execution"
	fmt.Println(minDistance(word1, word2))
}
