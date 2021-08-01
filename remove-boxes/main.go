package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func removeBoxes(boxes []int) int {
	if len(boxes) == 0 {
		return 0
	}
	var dp [101][101][101]int

	// base case
	for i := 0; i < len(boxes); i++ {
		for k := 0; k <= i; k++ {
			dp[i][i][k] = (k + 1) * (k + 1)
		}
	}

	for l := 1; l < len(boxes); l++ {
		for i := 0; i < len(boxes)-l; i++ {
			j := i + l

			for k := 0; k <= i; k++ {
				res := (k+1)*(k+1) + dp[i+1][j][0]

				for m := i + 1; m <= j; m++ {
					if boxes[m] == boxes[i] {
						res = max(res, dp[i+1][m-1][0]+dp[m][j][k+1])
					}
				}
				dp[i][j][k] = res
			}
		}
	}
	return dp[0][len(boxes)-1][0]
}

func main() {
	fmt.Println(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}))
}
