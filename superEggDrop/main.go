package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)

	dp[1] = make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[1][i] = i
	}

	for i := 2; i <= k; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if i > j {
				dp[i][j] = dp[i-1][j]
				continue
			}

			var maxE int = 1000000

			// if the egg is broken at floor j1
			for j1 := 1; j1 <= j; j1++ {

				temp := 1 + max(dp[i-1][j1-1], dp[i][j-j1])
				if temp < maxE {
					maxE = temp
				}
			}

			dp[i][j] = maxE
		}
	}
	for i := 1; i <= k; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(" ")
			fmt.Print(dp[i][j])
		}
		fmt.Println()
	}

	return dp[k][n]
}

func main() {
	k := 2
	n := 6
	fmt.Println(superEggDrop(k, n))
}
