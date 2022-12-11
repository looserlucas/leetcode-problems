package main

import "fmt"

func maxChunksToSorted(arr []int) int {
	var curMax, curMin = -1, 11
	var first, last = 0, 0
	var res = 0
	for i := 0; i < len(arr); i++ {
		last = i
		if arr[i] > curMax {
			curMax = arr[i]
		}
		if arr[i] < curMin {
			curMin = arr[i]
		}
		if curMin == first && curMax == last {
			curMax, curMin = -1, 11
			if i+1 < len(arr) {
				first, last = i+1, i+1
			}
			res++
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
