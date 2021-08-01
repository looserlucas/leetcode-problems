package main

import "fmt"

func splitArray(nums []int, m int) int {
	r := 0
	l := 0
	for i := 0; i < len(nums); i++ {
		r += nums[i]
		if nums[i] > l {
			l = nums[i]
		}
	}

	for l < r {
		mid := l + (r-l)/2
		if checkSplitable(mid, m, nums) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

// check if current value (sum) is able to be get by splitting this array, if no, then the sum is too small and we have to increase it, or else, we need to decrease it
func checkSplitable(sum int, m int, nums []int) bool {
	currentSum := 0
	counter := 1
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		if currentSum > sum {
			currentSum = nums[i]
			counter++
			if counter > m {
				return false
			}
		}
	}
	return true
}

func main() {
	nums := []int{7, 2, 5, 10, 8}
	m := 2
	fmt.Println(splitArray(nums, m))
}
