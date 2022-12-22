package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func findTheLongestSubstring(s string) int {
	var cMap []int = []int{1, 0, 0, 0, 2, 0, 0, 0, 4, 0, 0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 16, 0, 0, 0, 0, 0}
	mask := 0
	res := 0
	var m = make([]int, 32)
	for i := 0; i < 32; i++ {
		m[i] = -1
	}
	for i := 0; i < len(s); i++ {
		mask ^= cMap[int(s[i])-int('a')]
		if mask != 0 && m[mask] == -1 {
			m[mask] = i
		}
		res = max(res, i-m[mask])
	}
	return res
}

func main() {
	//fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
	fmt.Println(findTheLongestSubstring("leetcodeisgreat"))
}
