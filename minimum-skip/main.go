package main

import (
	"fmt"
	"math"
)

func min(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}

// dp[i][j] = til k
func minSkips(dist []int, speed int, hoursBefore int) int {
	var dp [2000][2000]float64

	var temp []int
	temp = append(temp, 0)
	dist = append(temp, dist...)

	dp[0][0] = 0
	for i := 1; i < len(dist); i++ {
		dp[i][0] = math.Ceil(float64(dp[i-1][0]) + float64(dist[i])/float64(speed))
		for j := 1; j <= i+1; j++ {
			if j > i {
				dp[i][j] = dp[i][j-1]
				continue
			}
			dp[i][j] = min(dp[i-1][j-1]+float64(dist[i])/float64(speed), math.Ceil(dp[i-1][j]+float64(dist[i])/float64(speed)))
		}
	}
	for i := 0; i < len(dist); i++ {
		for j := 0; j < len(dist); j++ {
			fmt.Print(dp[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

	for i := 0; i < len(dist); i++ {
		if dp[len(dist)-1][i] <= float64(hoursBefore) {
			return i
		}
	}
	return -1
}

func main() {
	dist := []int{1, 3, 2}
	speed := 4
	hoursBefore := 2
	fmt.Println(minSkips(dist, speed, hoursBefore))
}
