package main

import "fmt"

var MODN = int(1e9 + 7)

/*
var dp [][]int

func numberOfWays(startPos int, endPos int, k int) int {
	dp = nil
	dp = make([][]int, 3000)
	for i := 0; i < len(dp); i++ {
		var tmp []int
		for j := 0; j <= k; j++ {
			tmp = append(tmp, -1)
		}
		dp[i] = tmp
	}
	return dfs(startPos, endPos, k)
}

func dfs(startPos, endPos, k int) int {
	if startPos == endPos && k == 0 {
		return 1
	}
	if k == 0 {
		return 0
	}
	if dp[startPos+1000][k] != -1 {
		return dp[startPos+1000][k]
	}
	res := (dfs(startPos-1, endPos, k-1) + dfs(startPos+1, endPos, k-1)) % MODN
	dp[startPos+1000][k] = res
	return res
}
*/
func numberOfWays(startPos int, endPos int, k int) int {
	dp := make([][]int, k+1)
	for i := 0; i < len(dp); i++ {
		tmp := make([]int, 3000)
		dp[i] = tmp
	}
	startPos += 1000
	endPos += 1000
	dp[1][startPos+1] = 1
	dp[1][startPos-1] = 1
	for i := 2; i <= k; i++ {
		for j := 1; j < len(dp[i])-1; j++ {
			dp[i][j] = (dp[i-1][j+1] + dp[i-1][j-1]) % MODN
		}
	}
	return dp[k][endPos]
}

func main() {
	fmt.Println("here")
	fmt.Println(numberOfWays(467, 113, 1000))
}
