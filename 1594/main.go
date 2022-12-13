package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

var MODN = int(1e9 + 7)

func maxProductPath(a [][]int) int {
	n := len(a)
	m := len(a[0])
	maxNeg := make([][]int, n)
	maxPos := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp1 := make([]int, m)
		tmp2 := make([]int, m)
		maxNeg[i] = tmp1
		maxPos[i] = tmp2
	}
	if a[0][0] > 0 {
		maxPos[0][0] = a[0][0]
		maxNeg[0][0] = 1
	} else if a[0][0] < 0 {
		maxPos[0][0] = -1
		maxNeg[0][0] = a[0][0]
	} else {
		maxNeg[0][0] = 0
		maxPos[0][0] = 0
	}

	var move [][]int = [][]int{{-1, 0}, {0, -1}}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				continue
			}
			maxPos[i][j] = -1
			maxNeg[i][j] = 1
			for _, k := range move {
				if i+k[0] >= 0 && j+k[1] >= 0 {
					if a[i][j] > 0 {
						if maxPos[i+k[0]][j+k[1]] != -1 {
							maxPos[i][j] = max(maxPos[i][j], maxPos[i+k[0]][j+k[1]]*a[i][j])
						}
						if maxNeg[i+k[0]][j+k[1]] != 1 {
							maxNeg[i][j] = min(maxNeg[i][j], maxNeg[i+k[0]][j+k[1]]*a[i][j])
						}
					} else if a[i][j] < 0 {
						if maxNeg[i+k[0]][j+k[1]] != 1 {
							maxPos[i][j] = max(maxPos[i][j], maxNeg[i+k[0]][j+k[1]]*a[i][j])
						}
						if maxPos[i+k[0]][j+k[1]] != -1 {
							maxNeg[i][j] = min(maxNeg[i][j], maxPos[i+k[0]][j+k[1]]*a[i][j])
						}
					} else {
						maxNeg[i][j] = 0
						maxPos[i][j] = 0
					}
				}
			}
		}
	}
	/*
		print2DArray(a)

		print2DArray(maxNeg)

		print2DArray(maxPos)
	*/
	if maxPos[n-1][m-1] == -1 {
		return -1
	} else {
		return maxPos[n-1][m-1] % MODN
	}
}

func print2DArray(a [][]int) {
	n := len(a)
	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}
	fmt.Println()
}
func main() {
	/*
		fmt.Println(maxProductPath([][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}}))
		fmt.Println(maxProductPath([][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}))
		fmt.Println(maxProductPath([][]int{{1, 3}, {0, -4}}))
	*/

	fmt.Println(maxProductPath([][]int{{-1, -4, 2}, {4, 3, -1}, {2, -4, 4}, {1, -1, -4}}))
}
