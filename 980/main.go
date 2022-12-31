package main

import "fmt"

var res, n, m int
var move = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func uniquePathsIII(a [][]int) int {
	n = len(a)
	m = len(a[0])
	res = 0
	var is, js int
	var target int = 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] != -1 {
				if a[i][j] == 1 {
					is = i
					js = j
				}
				newMask := 1 << (i*m + j)
				target |= newMask
			}
		}
	}
	dfs(a, is, js, 0|(1<<(is*m+js)), target)
	return res
}

func dfs(a [][]int, i, j int, mask, target int) {
	if a[i][j] == 2 {
		res++
		return
	}
	for _, k := range move {
		x := i + k[0]
		y := j + k[1]
		if x >= 0 && x < n && y >= 0 && y < m && a[x][y] != -1 {
			newMask := 1 << (x*m + y)
			if mask&newMask == 0 {
				//fmt.Println(mask | newMask)
				if a[x][y] == 2 {
					if mask|newMask == target {
						dfs(a, x, y, mask|newMask, target)
					}
				} else {
					dfs(a, x, y, mask|newMask, target)
				}
			}
		}
	}
}

func main() {
	fmt.Println(uniquePathsIII([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}))
}
