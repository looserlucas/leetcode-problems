package main

import (
	"fmt"
	"sort"
)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	var sumMap1 = make(map[int]int)
	var sumMap2 = make(map[int]int)
	var sum1, sum2 [][]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sumMap1[nums1[i]+nums2[j]]++
			sumMap2[nums3[i]+nums4[j]]++
		}
	}
	for k, v := range sumMap1 {
		sum1 = append(sum1, []int{k, v})
	}
	for k, v := range sumMap2 {
		sum2 = append(sum2, []int{k, v})
	}
	sort.Slice(sum1, func(i, j int) bool {
		return sum1[i][0] < sum1[j][0]
	})
	sort.Slice(sum2, func(i, j int) bool {
		return sum2[i][0] < sum2[j][0]
	})
	res := 0
	for i := 0; i < len(sum1); i++ {
		t := -sum1[i][0]
		x := sort.Search(len(sum2), func(j int) bool {
			return sum2[j][0] >= t
		})
		if x == len(sum2) {
			continue
		} else {
			if sum2[x][0] == t {
				res += sum1[i][1] * sum2[x][1]
			}
		}
	}
	return res
}
func main() {
	fmt.Println("vim-go")
}
