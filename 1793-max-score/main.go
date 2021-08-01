package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// start with k i=k, j=k, if a[i-1] < a[j+1] => max result can get with (j-i+1) is when we move j++, else i--
func maximumScore(nums []int, k int) int {
	res := nums[k]
	mini := nums[k]
	i := k
	j := k
	for i > 0 || j < len(nums)-1 {
		if i == 0 {
			j++
		} else if j == len(nums)-1 {
			i--
		} else if nums[i-1] < nums[j+1] {
			j++
		} else {
			i--
		}
		mini = min(mini, min(nums[i], nums[j]))
		res = max(res, (j-i+1)*mini)
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
