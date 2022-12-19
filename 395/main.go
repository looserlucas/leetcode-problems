package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestSubstring(s string, k int) int {
	var seen = make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		seen[s[i]] = true
	}
	count := make(map[byte]int)
	maxUnique := len(seen)
	var res = 0
	for u := 1; u <= maxUnique; u++ {
		j := 0
		for i := 0; i < len(s); i++ {
			count[s[i]]++
			for j < i && len(count) > u {
				count[s[j]]--
				if count[s[j]] == 0 {
					delete(count, s[j])
				}
				j++
			}
			ok := true
			for _, v := range count {
				if v < k {
					ok = false
				}
			}
			if ok {
				res = max(res, i-j+1)
			}
		}
		for k, _ := range count {
			delete(count, k)
		}
	}
	return res
}

func main() {
	fmt.Println(longestSubstring("aaabb", 3))
	fmt.Println(longestSubstring("ababbc", 2))
}
