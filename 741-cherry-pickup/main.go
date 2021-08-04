package main

import "fmt"

const MIN_INT = -9999999

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func cherryPickup(grid [][]int) int {
	var memo [51][51][51]int

	for i := 0; i < 51; i++ {
		for j := 0; j < 51; j++ {
			for k := 0; k < 51; k++ {
				memo[i][j][k] = MIN_INT
			}
		}
	}
	ans := cherryPickupMemo(grid, &memo, 0, 0, 0)
	return max(0, ans)
}

func cherryPickupMemo(grid [][]int, memo *[51][51][51]int, r1, c1, c2 int) int {
	r2 := c1 + r1 - c2

	if r1 == len(grid) || r2 == len(grid) || c1 == len(grid) || c2 == len(grid) || grid[r1][c1] == -1 || grid[r2][c2] == -1 {
		return MIN_INT
	} else if c1 == len(grid)-1 && r1 == len(grid)-1 {
		return grid[r1][c1]
	} else if memo[r1][c1][c2] != MIN_INT {
		return memo[r1][c1][c2]
	} else {
		ans := grid[r1][c1]
		// if c1 != c2 => r1 != r2
		if c1 != c2 {
			ans += grid[r2][c2]
		}
		ans += max(max(cherryPickupMemo(grid, memo, r1+1, c1, c2+1), cherryPickupMemo(grid, memo, r1+1, c1, c2)), max(cherryPickupMemo(grid, memo, r1, c1+1, c2), cherryPickupMemo(grid, memo, r1, c1+1, c2+1)))
		memo[r1][c1][c2] = ans
		return ans
	}
}

func main() {
	grid := [][]int{{1, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 1}, {1, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 1, 1, 1}}
	fmt.Println(cherryPickup(grid))
}
