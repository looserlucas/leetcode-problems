package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// var dp[k][i] = min space wasted from 1->i with k cut
// dp[k][i] = min(dp[k][i], dp[k-1][j]+ max(j+1 -> i) * (j-i) - sum(nums[j+1], nums[i])))
func minSpaceWastedKResizing(nums []int, k int) int {
	var dp [][]int = make([][]int, k+1)
	var curMax, sum = 0, 0
	var tmp []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > curMax {
			curMax = nums[i]
		}
		sum += nums[i]
		tmp = append(tmp, curMax*(i+1)-sum)
	}
	dp[0] = tmp
	for i := 1; i <= k; i++ {
		var tmp = make([]int, len(nums))
		dp[i] = tmp
	}

	for t := 1; t <= k; t++ {
		for i := 0; i < len(nums); i++ {
			curMax := 0
			sum := 0
			dp[t][i] = dp[t-1][i]
			for j := 1; j <= i; j++ {
				if nums[i-j+1] > curMax {
					curMax = nums[i-j+1]
				}
				sum += nums[i-j+1]
				dp[t][i] = min(dp[t][i], dp[t-1][i-j]+curMax*j-sum)
			}
		}
	}
	return dp[k][len(nums)-1]
}

func main() {
	fmt.Println(minSpaceWastedKResizing([]int{10, 20}, 0))
	fmt.Println(minSpaceWastedKResizing([]int{10, 20, 30}, 1))
	fmt.Println(minSpaceWastedKResizing([]int{10, 20, 15, 30, 20}, 2))
}
