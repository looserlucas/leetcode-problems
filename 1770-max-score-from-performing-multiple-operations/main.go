package main

import "fmt"

const MIN_INT = -int(1e9)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maximumScore(nums []int, multipliers []int) int {
	var memo [][]int
	for i := 0; i <= len(multipliers); i++ {
		var tmp []int
		for j := 0; j <= len(multipliers); j++ {
			tmp = append(tmp, MIN_INT)
		}
		memo = append(memo, tmp)
	}

	return maximumScoreTopDown(nums, multipliers, &memo, 0, 0)
}

func maximumScoreTopDown(nums []int, multipliers []int, memo *[][]int, l, i int) int {
	if i >= len(multipliers) {
		return 0
	}
	if (*memo)[l][i] == MIN_INT {
		r := len(nums) - 1 - (i - l)
		(*memo)[l][i] = max(nums[l]*multipliers[i]+maximumScoreTopDown(nums, multipliers, memo, l+1, i+1), nums[r]*multipliers[i]+maximumScoreTopDown(nums, multipliers, memo, l, i+1))
	}
	return (*memo)[l][i]
}

func main() {
	nums := []int{-5, -3, -3, -2, 7, 1}
	multipliers := []int{-10, -5, 3, 4, 6}
	fmt.Println(maximumScore(nums, multipliers))
}
