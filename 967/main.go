package main

import "fmt"

var res []int

func numsSameConsecDiff(n int, k int) []int {
	res = nil
	for i := 1; i <= 9; i++ {
		dfs(i, 1, n, k)
	}
	return res
}

func dfs(now, count, n, k int) {
	if count == n {
		res = append(res, now)
		return
	}
	last := now % 10
	if last-k == last+k {
		dfs(now*10+(last-k), count+1, n, k)
	} else {
		if last-k >= 0 {
			dfs(now*10+(last-k), count+1, n, k)
		}
		if last+k <= 9 {
			dfs(now*10+(last+k), count+1, n, k)
		}
	}
}

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
	fmt.Println(numsSameConsecDiff(2, 1))
	fmt.Println(numsSameConsecDiff(3, 0))
}
