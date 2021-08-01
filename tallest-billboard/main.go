package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func tallestBillboard(rods []int) int {
	var memo = make([]map[int]int, len(rods))
	for i := 0; i < len(rods); i++ {
		tmp := make(map[int]int)
		memo[i] = tmp
	}
	return dfsWithMemo(rods, 0, 0, 0, memo)
}

// dfs with memo
// memo[i][abs(s1-s2)] = the max additional value of the billboard that we can get from (i+1, n) with sum1=s1, sum2=s2
// memo[i][abs(s1-s2)] = -1 if from i+1,n we can't create a billboard
func dfsWithMemo(rods []int, cur, sum1, sum2 int, memo []map[int]int) int {
	if cur == len(rods) {
		if sum1 == sum2 {
			return sum1
		} else {
			return -1
		}
	}

	if _, ok := memo[cur][abs(sum1, sum2)]; !ok {
		memo[cur][abs(sum1, sum2)] = max(-1, max(dfsWithMemo(rods, cur+1, sum1, sum2, memo), max(dfsWithMemo(rods, cur+1, sum1+rods[cur], sum2, memo), dfsWithMemo(rods, cur+1, sum1, sum2+rods[cur], memo)))-max(sum1, sum2))
	}

	if memo[cur][abs(sum1, sum2)] == -1 {
		return memo[cur][abs(sum1, sum2)]
	}

	return memo[cur][abs(sum1, sum2)] + max(sum1, sum2)
}

// naive dfs before memo
func dfs(rods []int, cur, sum1, sum2 int, res *int) {
	if cur == len(rods) {
		if sum1 == sum2 {
			*res = max(*res, sum1)
		}
		return
	}
	dfs(rods, cur+1, sum1, sum2, res)
	dfs(rods, cur+1, sum1+rods[cur], sum2, res)
	dfs(rods, cur+1, sum1, sum2+rods[cur], res)
	return
}

func main() {
	rods := []int{981, 472, 813, 534, 616, 854, 265, 402, 30, 20, 3, 4}
	fmt.Println(tallestBillboard(rods))
}
