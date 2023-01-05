package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findMaxLength(a []int) int {
	var sum = make([]int, len(a)+1)
	var seen = make(map[int]int)
	sum[0] = 0
	seen[0] = 0
	var res int = 0
	for i := 0; i < len(a); i++ {
		if a[i] == 1 {
			sum[i+1] = sum[i] + 1
		} else {
			sum[i+1] = sum[i] - 1
		}
		v, ok := seen[sum[i+1]]
		if !ok {
			seen[sum[i+1]] = i + 1
		} else {
			res = max(res, i+1-v)
		}
	}
	return res
}

func main() {
	fmt.Println(findMaxLength([]int{1}))
	fmt.Println(findMaxLength([]int{0, 1, 0, 1, 0, 0, 1}))
}
