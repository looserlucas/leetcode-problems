package main

import "fmt"

func combinationSum4(n []int, target int) int {
	memo := make([]int, target+1)
	memo[0] = 0
	for i := 1; i <= target; i++ {
		memo[i] = 0
	}
	for j := 0; j < len(n); j++ {
		if n[j] <= target {
			memo[n[j]] = 1
		}
	}

	for i := 1; i <= target; i++ {
		for j := 0; j < len(n); j++ {
			if i-n[j] > 0 {
				fmt.Println(i, i-n[j])
				memo[i] = memo[i] + memo[i-n[j]]
			}
		}
	}
	return memo[target]
}

func main() {
	fmt.Println(combinationSum4([]int{9}, 3))
}
