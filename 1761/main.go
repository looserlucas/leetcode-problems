package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func minTrioDegree(n int, edges [][]int) int {
	var degree = make([]int, n+1)
	var seen = make([][]bool, n+1)
	for i := 0; i < len(seen); i++ {
		var tmp = make([]bool, n+1)
		seen[i] = tmp
	}
	for i := 0; i < len(edges); i++ {
		seen[edges[i][0]][edges[i][1]] = true
		seen[edges[i][1]][edges[i][0]] = true
		degree[edges[i][0]]++
		degree[edges[i][1]]++
	}
	var res int = math.MaxInt
	for i := 1; i <= n-2; i++ {
		for j := i + 1; j <= n-1; j++ {
			for k := j + 1; k <= n; k++ {
				if seen[i][j] && seen[i][k] && seen[j][k] {
					res = min(res, degree[i]+degree[j]+degree[k]-6)
				}
			}
		}
	}
	if res == math.MaxInt {
		return -1
	} else {
		return res
	}
}

func main() {
	//fmt.Println(minTrioDegree(6, [][]int{{1, 2}, {1, 3}, {3, 2}, {4, 1}, {5, 2}, {3, 6}}))
	fmt.Println(minTrioDegree(7, [][]int{{1, 3}, {4, 1}, {4, 3}, {2, 5}, {5, 6}, {6, 7}, {7, 5}, {2, 6}}))
}
