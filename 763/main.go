package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func partitionLabels(s string) []int {
	st := make(map[byte]int)
	en := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := st[s[i]]; !ok {
			st[s[i]] = i
		}
	}

	for j := len(s) - 1; j >= 0; j-- {
		if _, ok := en[s[j]]; !ok {
			en[s[j]] = j
		}
	}

	var a [][]int
	for k, v := range st {
		a = append(a, []int{v, en[k]})
	}

	sort.SliceStable(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})

	var res []int
	startNow := -1
	endNow := -1
	for i := 0; i < len(a); i++ {
		if startNow == -1 {
			startNow = a[i][0]
			endNow = a[i][1]
			continue
		}
		if a[i][0] < endNow {
			endNow = max(endNow, a[i][1])
		} else {
			res = append(res, endNow-startNow+1)
			startNow = a[i][0]
			endNow = a[i][1]
		}
		if i+1 == len(a) {
			res = append(res, endNow-startNow+1)
		}
	}
	return res
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}
