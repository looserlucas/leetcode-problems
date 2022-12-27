package main

import "fmt"

func canJump(nums []int) bool {
	curMax := 0
	for i := 0; i < len(nums); i++ {
		if i > curMax {
			return false
		}
		if i+nums[i] > curMax {
			curMax = i + nums[i]
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
}
