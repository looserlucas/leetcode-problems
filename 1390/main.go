package main

import (
	"fmt"
)

func sumFourDivisors(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		cnt := 0
		for j := 1; j*j <= nums[i]; j++ {
			if nums[i]%j == 0 {
				if nums[i]/j == j {
					sum += j
					cnt++
				} else {
					sum += j + nums[i]/j
					cnt += 2
				}
			}
		}
		if cnt == 4 {
			res += sum
		}
	}
	return res
}
func main() {
	fmt.Println(sumFourDivisors([]int{21, 4, 7}))
	fmt.Println(sumFourDivisors([]int{21, 21}))
	fmt.Println(sumFourDivisors([]int{1, 2, 3, 4, 5}))
}
