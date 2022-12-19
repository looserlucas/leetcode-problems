package main

import "fmt"

func atMost(nums []int, k int) int {
	j := 0
	nowOdd := 0
	var res = 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			nowOdd++
		}
		for j <= i && nowOdd > k {
			if nums[j]%2 != 0 {
				nowOdd--
			}
			j++
		}
		fmt.Println(i, j)
		res += i - j + 1
	}
	return res
}

func numberOfSubarrays(nums []int, k int) int {
	return atMost(nums, k) - atMost(nums, k-1)
}

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 1, 1, 1}, 1))
}
