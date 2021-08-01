package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProduct(nums []int) int {
	maxNeg := 1
	pro := 1
	res := -99999999

	for i := 0; i < len(nums); i++ {
		pro *= nums[i]

		if pro >= 0 {
			res = max(res, pro)
			if pro == 0 {
				maxNeg = 1
				pro = 1
			}
		} else {
			if maxNeg == 1 {
				maxNeg = pro
				res = max(res, pro)
				continue
			} else {
				res = max(res, pro/maxNeg)
				if maxNeg < pro {
					maxNeg = pro
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{0, -1, 0, 0, -1, 0, -1, 0, -1, 0}
	fmt.Println(maxProduct(nums))
}
