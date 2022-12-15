package main

import "fmt"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	nowRow := make([]float64, 1)
	nowRow[0] = float64(poured)
	for i := 0; i <= query_row; i++ {
		newRow := make([]float64, len(nowRow)+2)
		for j := 0; j < len(nowRow); j++ {
			if nowRow[j] > 1 {
				newRow[j] += (nowRow[j] - 1) / 2
				newRow[j+1] += (nowRow[j] - 1) / 2
				nowRow[j] = 1
			}
		}
		if i != query_row {
			nowRow = newRow
		}
	}
	return nowRow[query_glass]
}

func main() {
	fmt.Println(champagneTower(1, 1, 1))
	fmt.Println(champagneTower(2, 1, 1))
	fmt.Println(champagneTower(100000009, 33, 17))
}
