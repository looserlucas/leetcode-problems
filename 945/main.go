package main

import "fmt"

func minIncrementForUnique(nums []int) int {
	bucket := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		bucket[nums[i]]++
	}
	res := 0
	for i := 0; i <= int(2e5); i++ {
		v, ok := bucket[i]
		if ok {
			if v > 1 {
				bucket[i] = 1
				res += v - 1
				bucket[i+1] += v - 1
			}
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
