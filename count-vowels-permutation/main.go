package main

import "fmt"

var dp [][]int

const moder = 1000000007

// we turn the vowel into number as follow:
// a-0, e-1, i-2, o-3, u-4

/*
This means, we have these state transformation:
a -> e, 0 -> 1
e -> a, i, 1 -> 0, 1 -> 2
i -> ..., 2 -> 0, 2 -> 1, 2 -> 3, 2 -> 4
o -> i, u, 3 -> 2, 3 -> 4
u -> a, 4 -> 0

We then define dp[j][i] as the number of mutation that start with character j, and has the length of i. The formula is as follow:
			if j == 0 {
				dp[j][i] = (dp[j][i] + dp[j+1][i-1]) % moder
			}
			if j == 1 {
				dp[j][i] = (dp[j][i] + dp[j-1][i-1] + dp[j+1][i-1]) % moder
			}
			if j == 2 {
				dp[j][i] = (dp[j][i] + dp[j-2][i-1] + dp[j-1][i-1] + dp[j+1][i-1] + dp[j+2][i-1]) % moder
			}
			if j == 3 {
				dp[j][i] = (dp[j][i] + dp[j-1][i-1] + dp[j+1][i-1]) % moder
			}
			if j == 4 {
				dp[j][i] = (dp[j][i] + dp[j-4][i-1]) % moder
			}

*/
func main() {
	var n, res int
	var dp [][]int
	n = 5
	n = n - 1
	for i := 0; i < 5; i++ {
		var temp []int
		temp = append(temp, 1)
		for j := 1; j <= n; j++ {
			temp = append(temp, 0)
		}
		dp = append(dp, temp)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < 5; j++ {
			if j == 0 {
				dp[j][i] = (dp[j][i] + dp[j+1][i-1]) % moder
			}
			if j == 1 {
				dp[j][i] = (dp[j][i] + dp[j-1][i-1] + dp[j+1][i-1]) % moder
			}
			if j == 2 {
				dp[j][i] = (dp[j][i] + dp[j-2][i-1] + dp[j-1][i-1] + dp[j+1][i-1] + dp[j+2][i-1]) % moder
			}
			if j == 3 {
				dp[j][i] = (dp[j][i] + dp[j-1][i-1] + dp[j+1][i-1]) % moder
			}
			if j == 4 {
				dp[j][i] = (dp[j][i] + dp[j-4][i-1]) % moder
			}
		}
	}

	fmt.Println(dp)
	for i := 0; i < 5; i++ {
		res = (res + dp[i][n]) % moder
	}

	fmt.Println(res)
}
