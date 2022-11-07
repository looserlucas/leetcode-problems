package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring(s string) int {
	i := 0
	j := 0
	var res int = 0
	var cur int = 0
	seen := make(map[byte]int)
	for j < len(s) {
		v, ok := seen[s[j]]
		if !ok {
			seen[s[j]] = j
			cur++
		} else {
			if i < v+1 {
				i = v + 1
			}
			cur = j - i + 1
			seen[s[j]] = j
		}
		res = max(res, cur)
		fmt.Println(i, j, cur)
		j++
	}
	return res
}

func main() {
	/*
			fmt.Println(lengthOfLongestSubstring("abcabcbb"))
			fmt.Println(lengthOfLongestSubstring("bbbbb"))
		fmt.Println(lengthOfLongestSubstring("abba"))
	*/
	fmt.Println(lengthOfLongestSubstring("abba"))
}
