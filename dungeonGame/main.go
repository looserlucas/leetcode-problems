package main

import "fmt"

func calculateMinimumHP(dungeon [][]int) int {
	var dp [1000][1000]int
	n := len(dungeon)
	m := len(dungeon[0])
	var temp int
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i == n-1 && j == m-1 {
				dp[i][j] = dungeon[i][j]
				continue
			}
			if j == m-1 && i != n-1 {
				temp = dp[i+1][j]
				if dungeon[i][j] < 0 {
					if temp < 0 {
						dp[i][j] = dungeon[i][j] + temp
					}
					dp[i][j] = dungeon[i][j]
				} else {
					dp[i][j] = temp + dungeon[i][j]
				}
				continue
			}
			if i == n-1 && j != m-1 {
				temp = dp[i][j+1]
				if dungeon[i][j] < 0 {
					if temp < 0 {
						dp[i][j] = dungeon[i][j] + temp
					}
					dp[i][j] = dungeon[i][j]
				} else {
					dp[i][j] = temp + dungeon[i][j]
				}
				continue
			}
			if dp[i+1][j] > dp[i][j+1] {
				temp = dp[i+1][j]
			} else {
				temp = dp[i][j+1]
			}

			if dungeon[i][j] < 0 {
				if temp < 0 {
					dp[i][j] = dungeon[i][j] + temp
					continue
				}
				dp[i][j] = dungeon[i][j]
			} else {
				dp[i][j] = temp + dungeon[i][j]
			}

		}
	}
	return -dp[0][0] + 1
}

func main() {
	dungeon := [][]int{{0, -3}}
	fmt.Println(calculateMinimumHP(dungeon))
}
