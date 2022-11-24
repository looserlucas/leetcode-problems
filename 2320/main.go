package main

import "fmt"

var MODN = int(1e9 + 7)

func countHousePlacements(n int) int {
	var dp [][4]int = make([][4]int, n+1)
	dp[1][0] = 1
	dp[1][1] = 1
	dp[1][2] = 1
	dp[1][3] = 1
	for i := 2; i <= n; i++ {
		dp[i][0] = (dp[i-1][3] + dp[i-1][2] + dp[i-1][1] + dp[i-1][0]) % MODN
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % MODN
		dp[i][2] = (dp[i-1][0] + dp[i-1][1]) % MODN
		dp[i][3] = (dp[i-1][0]) % MODN
	}
	fmt.Println(dp)
	return (dp[n][1] + dp[n][2] + dp[n][3] + dp[n][0]) % MODN
}
func main() {
	fmt.Println(countHousePlacements(3))
}
