package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestNiceSubarray(nums []int) int {
	used := 0
	j := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		for used&nums[i] != 0 {
			used ^= nums[j]
			j++
		}
		used |= nums[i]
		res = max(res, i-j+1)
	}

	return res
}

func main() {
	fmt.Println(longestNiceSubarray([]int{84139415, 693324769, 614626365, 497710833, 615598711, 264, 65552, 50331652, 1, 1048576, 16384, 544, 270532608, 151813349, 221976871, 678178917, 845710321, 751376227, 331656525, 739558112, 267703680}))
}
