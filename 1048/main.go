package main

import (
	"fmt"
	"sort"
)

func isWordChange(a, b string) bool {
	if len(b) != len(a)+1 {
		return false
	}
	j := 0
	i := 0
	ok := false
	for i < len(a) {
		if b[j] != a[i] {
			if !ok {
				ok = true
				j++
				continue
			} else {
				return false
			}
		}
		i++
		j++
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func less(a, b string) bool {
	if len(a) < len(b) {
		return true
	} else if len(a) > len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return false
		} else if a[i] < b[i] {
			return true
		}
	}
	return false
}

func longestStrChain(words []string) int {
	var dp []int = make([]int, len(words))
	sort.Slice(words, func(i, j int) bool {
		return less(words[i], words[j])
	})
	var res int = 0
	for i := 0; i < len(words); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if isWordChange(words[j], words[i]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func main() {
	//fmt.Println(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
	fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
}
