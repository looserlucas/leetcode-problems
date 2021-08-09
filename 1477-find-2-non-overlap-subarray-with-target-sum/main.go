package main

import "fmt"

const MAX_INT = int(1e9)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func minSumOfLengths(arr []int, target int) int {
	var dd = make(map[int]int)
	dd[0] = -1
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		dd[sum] = i
	}

	sum = 0
	lsize := MAX_INT
	res := MAX_INT
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if v, ok := dd[sum-target]; ok {
			lsize = min(lsize, i-v)
		}

		if v, ok := dd[sum+target]; ok && lsize < MAX_INT {
			res = min(res, lsize+v-i)
		}
	}

	if res != MAX_INT {
		return res
	} else {
		return -1
	}
}

func main() {
	arr := []int{4, 3, 2, 6, 2, 3, 4}
	target := 6
	fmt.Println(minSumOfLengths(arr, target))
}
