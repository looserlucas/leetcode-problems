package main

import "fmt"

const MAX_NUM = int(1e5 + 1)
const MODER = int(1e9 + 7)

// Explaination: use freq to store the frequency of the number in nums, then we turn freq into a prefix sum of it self -> freq turns into frequency prefix sum. This means if we want to find out how many numbers in array between [l,r], it s freq[r]-freq[l-1]
// this solution is a better O(n^2) solution than a 2 pointer one, but still gonna give TLE
func sumOfFlooredPairs(nums []int) int {
	var freq [2*MAX_NUM + 1]int
	max := 0
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
		if nums[i] > max {
			max = nums[i]
		}
	}
	for i := 1; i < 2*MAX_NUM+1; i++ {
		freq[i] += freq[i-1]
	}

	var sum int
	for i := 0; i < len(nums); i++ {
		l := nums[i]
		r := 2*nums[i] - 1
		div := 1
		for l <= max {
			sum = (sum + div*(freq[r]-freq[l-1])) % MODER
			l += nums[i]
			r += nums[i]
			div++
		}
	}
	return sum
}

func main() {
	nums := []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Println(sumOfFlooredPairs(nums))
}
