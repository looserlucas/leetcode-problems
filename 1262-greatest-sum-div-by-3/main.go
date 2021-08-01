package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

/*
dp[0] := max sum % 3 == 0
dp[1] := max sum % 3 == 1
dp[2] := max sum % 3 == 2
*/
func maxSumDivThree(nums []int) int {
	var dp [3]int

	for i := 0; i < len(nums); i++ {
		r := dp
		for j := 0; j < 3; j++ {
			dp[(nums[i]+r[j])%3] = max(dp[(nums[i]+r[j])%3], nums[i]+r[j])
		}
		fmt.Println(dp)
	}
	return dp[0]
}

func main() {
	nums := []int{1, 2, 3, 4, 4}
	fmt.Println(maxSumDivThree(nums))
}
