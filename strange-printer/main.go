package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// dp[i][j] -> the minimum way to print from s[i] -> s[j-1]
// => dp[i][j] = d[i][k] + dp[k][j]
// O(n^3) since we need 1 for loop for i, 1 for loop for j, and 1 more for k
func strangePrinter(s string) int {
	var dp [101][101]int
	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
	}

	// first way
	/*
		for i := len(s) - 1; i >= 0; i-- {
			for dist := 1; dist+i < len(s); dist++ {
				j := dist + i
				if dist == 1 {
					if s[i] == s[j] {
						dp[i][j] = 1
					} else {
						dp[i][j] = 2
					}
				}
				dp[i][j] = 1000
				for k := i; k < j; k++ {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
				}
				if s[i] == s[j] {
					dp[i][j]--
				}
			}
		}
	*/

	// second way
	for dist := 1; dist < len(s); dist++ {
		for left := 0; left < len(s)-dist; left++ {
			right := left + dist
			if dist == 1 {
				if s[left] == s[right] {
					dp[left][right] = 1
					continue
				} else {
					dp[left][right] = 2
					continue
				}
			}
			dp[left][right] = 1000
			for k := left; k < right; k++ {
				dp[left][right] = min(dp[left][right], dp[left][k]+dp[k+1][right])
			}
			if s[right] == s[left] {
				dp[left][right]--
			}
		}
	}
	return dp[0][len(s)-1]
}

func main() {
	fmt.Println(strangePrinter("aaabbb"))
}
