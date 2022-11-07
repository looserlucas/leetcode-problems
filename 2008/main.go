package main

import (
	"fmt"
	"sort"
)

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

var memo []int64
var r [][]int

func maxTaxiEarnings(n int, rides [][]int) int64 {
	memo = nil
	for i := 0; i < len(rides); i++ {
		memo = append(memo, -1)
	}
	sort.SliceStable(rides, func(i, j int) bool {
		if rides[i][0] < rides[j][0] {
			return true
		} else if rides[i][0] == rides[j][0] {
			if rides[i][1] < rides[j][1] {
				return true
			} else if rides[i][1] == rides[j][1] {
				if rides[i][2] > rides[j][2] {
					return true
				}
			}
		}
		return false
	})
	r = rides

	return dfs(0)
}

func dfs(index int) int64 {
	if len(r) == index {
		return 0
	}
	if memo[index] != -1 {
		return memo[index]
	}

	end := r[index][1]
	next := sort.Search(len(r), func(i int) bool {
		return r[i][0] >= end
	})
	res := max(dfs(index+1), dfs(next)+int64(r[index][1])-int64(r[index][0])+int64(r[index][2]))
	memo[index] = res
	return res
}

func main() {
	fmt.Println(maxTaxiEarnings(20, [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}}))
	fmt.Println(maxTaxiEarnings(20, [][]int{{2, 5, 4}, {1, 5, 1}}))
}
