package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// dp[i][j][k] = minimum way to turn stones[i] -> stone[j] into k piles
// dp[i][j][1] = dp[i][j][k] + sum[i][j]
// minimum ways to turn i->j stone to k pile = min ways to turn i-> j1 to k-1 pile and j1+1 -> j to 1 pile
// => dp[i][j][k] = min(dp[i][j][k], dp[i][j1][k-1] + dp[j1+1][j][1]) with j1 in [i,j]
func mergeStones(stones []int, k int) int {
	var prefix []int = make([]int, len(stones))
	var dp [40][40][40]int

	if (len(stones)-1)%(k-1) != 0 {
		return -1
	}

	var max = 99999999
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(stones); j++ {
			for k1 := 1; k1 <= k; k1++ {
				dp[i][j][k1] = max
			}
		}
	}

	for i := 0; i < len(stones); i++ {
		dp[i][i][1] = 0
	}

	prefix[0] = stones[0]
	for i := 1; i < len(stones); i++ {
		prefix[i] = prefix[i-1] + stones[i]
	}

	for dist := 1; dist < len(stones); dist++ {
		for i := 0; i < len(stones)-dist; i++ {
			j := i + dist
			for k1 := 2; k1 <= k; k1++ {
				for j1 := i; j1 <= j; j1++ {
					dp[i][j][k1] = min(dp[i][j][k1], dp[i][j1][k1-1]+dp[j1+1][j][1])
				}
			}
			dp[i][j][1] = dp[i][j][k] + prefix[j] - prefix[i] + stones[i]
		}
	}

	return dp[0][len(stones)-1][1]
}

func main() {
	stones := []int{3, 2, 4, 1}
	k := 3
	fmt.Println(mergeStones(stones, k))
}
