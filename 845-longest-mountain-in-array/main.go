package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func longestMountain(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var dd1 = make([]int, len(arr))
	var dd2 = make([]int, len(arr))
	dd1[0] = 1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			dd1[i] = dd1[i-1] + 1
		} else {
			dd1[i] = 1
		}
	}
	dd2[len(arr)-1] = 1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			dd2[i] = dd2[i+1] + 1
		} else {
			dd2[i] = 1
		}
	}

	res := 0
	for i := 0; i < len(arr); i++ {
		if dd1[i] == 1 || dd2[i] == 1 {
			continue
		}
		if dd1[i]+dd2[i]-1 >= 3 {
			res = max(res, dd1[i]+dd2[i]-1)
		}
	}
	return res
}
func main() {
	arr := []int{2, 2, 2}
	fmt.Println(longestMountain(arr))
}
