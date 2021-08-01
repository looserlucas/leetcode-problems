package main

import "fmt"

func checkPossibility(nums []int) bool {
	count := 0
	i := 0
	for i < len(nums)-1 {
		if count >= 2 {
			return false
		}
		if i == 0 {
			if nums[i] > nums[i+1] {
				count++
			}
			i++
			continue
		} else {
			if nums[i] > nums[i+1] {
				count++
				if nums[i+1] < nums[i-1] {
					nums[i+1] = nums[i]
				}
			}
			i++
		}
	}
	if count >= 2 {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println("vim-go")
}
