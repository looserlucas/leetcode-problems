package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	i := 0
	j := len(matrix) - 1
	for i < j {
		mid := i + (j-i+1)/2
		if matrix[mid][0] > target {
			j = mid - 1
		} else {
			i = mid
		}
	}
	if i >= len(matrix) {
		i = len(matrix) - 1
	}

	row := i
	i = 0
	j = len(matrix[row]) - 1
	for i < j {
		mid := i + (j-i+1)/2
		if matrix[row][mid] > target {
			j = mid - 1
		} else {
			i = mid
		}
	}
	if i >= len(matrix[row]) {
		return false
	}
	if matrix[row][i] == target {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 64}}, 11))
	//fmt.Println(searchMatrix([][]int{{1}, {1}}, 0))
}
