package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var pre = make([]int, len(nums)+1)
	pre[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] + nums[i]
	}

	i := 0
	j := 0
	res := 999999
	for {
		if j == len(nums)-1 {
			if i == len(nums)-1 {
				break
			}
			if pre[j]-pre[i] >= target {
				res = min(res, j-i)
			}
			if i == 0 {
				if pre[j]-pre[i]+nums[i] >= target {
					res = min(res, j-i+1)
				}
			}
			i++
			continue
		}

		if pre[j]-pre[i] >= target {
			res = min(res, j-i)
			if i+1 < len(nums) && pre[j]-pre[i+1] >= target {
				i++
			} else {
				j++
			}
		} else {
			if i == 0 {
				if pre[j]-pre[i]+nums[i] >= target {
					res = min(res, j-i+1)
				}
			}
			j++
		}
	}

	if res == 999999 {
		return 0
	} else {
		return res
	}
}
func main() {
	target := 15
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(minSubArrayLen(target, nums))
}
