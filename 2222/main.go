package main

import "fmt"

func numberOfWays(s string) int64 {
	var dp [1e5 + 1][4][2]int64
	if s[0] == '1' {
		dp[0][1][int('1')-int('0')] = 1
	} else {
		dp[0][1][int('0')-int('0')] = 1
	}
	for i := 1; i < len(s); i++ {
		dp[i][1][int(s[i])-int('0')] = 1
		for j := 1; j <= 3; j++ {
			if s[i] == '1' {
				dp[i][j][int('1')-int('0')] += dp[i-1][j-1][int('0')-int('0')] + dp[i-1][j][int('1')-int('0')]
				dp[i][j][int('0')-int('0')] += dp[i-1][j][int('0')-int('0')]
			} else {
				dp[i][j][int('0')-int('0')] += dp[i-1][j-1][int('1')-int('0')] + dp[i-1][j][int('0')-int('0')]
				dp[i][j][int('1')-int('0')] += dp[i-1][j][int('1')-int('0')]
			}
		}
	}
	return dp[len(s)-1][3][int('0')-int('0')] + dp[len(s)-1][3][int('1')-int('0')]
}

/*
func dfs(s string, i int, k int, end int) int64 {
	if k == 0 {
		return 0
	}
	if i == 0 && k == 1 && end == int(s[i]) {
		return 1
	}
	if i == 0 {
		return 0
	}
	if dp[i][k][end] != -1 {
		return dp[i][k][end]
	}
	var res int64
	if s[i] == '1' {
		res = dfs(s, i-1, k-1, int('0')) + dfs(s, i-1, k, int('1'))
	} else {
		res = dfs(s, i-1, k-1, int('1')) + dfs(s, i-1, k, int('0'))
	}
	dp[i][k][end] = res
	return res
}
*/
func main() {
	fmt.Println(numberOfWays("001101"))
	fmt.Println(numberOfWays("11100"))
}
