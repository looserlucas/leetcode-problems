package main

import "fmt"

func canCross(stones []int) bool {
	if len(stones) <= 1 {
		return true
	}
	if stones[0]+1 != stones[1] {
		return false
	}

	var dp = make(map[int]map[int]bool)
	var tmp = make(map[int]bool)
	dp[0] = tmp
	tmp[1] = true
	dp[1] = tmp

	for i := 2; i < len(stones); i++ {
		dp[i] = make(map[int]bool)
		for j := 1; j < i; j++ {
			diff := stones[i] - stones[j]
			_, ok1 := dp[j][diff+1]
			_, ok2 := dp[j][diff]
			_, ok3 := dp[j][diff-1]
			if ok1 || ok2 || ok3 {
				dp[i][diff] = true
			}
		}
	}

	if val, ok := dp[len(stones)-1]; ok && len(val) != 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var stones = []int{0, 1, 2, 5, 6, 9, 10, 12, 13, 14, 17, 19, 20, 21, 26, 27, 28, 29, 30}
	fmt.Println(canCross(stones))
}
