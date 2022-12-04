package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

var MODN = int(1e9 + 7)

func maximumProduct(nums []int, k int) int {
	var bucket = make([]int, int(1e6+2e5))
	for i := 0; i < len(nums); i++ {
		bucket[nums[i]]++
	}
	for i := 0; i < len(bucket); i++ {
		if k > 0 && bucket[i] > 0 {
			x := min(bucket[i], k)
			bucket[i] -= x
			bucket[i+1] += x
			k -= x
		}
	}
	var res int = 1
	for i := 0; i < len(bucket); i++ {
		if bucket[i] != 0 {
			for j := 0; j < bucket[i]; j++ {
				res = (res % MODN * i % MODN) % MODN
			}
		}
	}
	return res
}

func main() {
	fmt.Println(maximumProduct([]int{0, 4}, 5))
	fmt.Println(maximumProduct([]int{6, 3, 3, 2}, 2))
}
