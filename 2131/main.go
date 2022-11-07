package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func reverse(s string) string {
	var res string
	res = res + string(s[1]) + string(s[0])
	return res
}

func longestPalindrome(words []string) int {
	var nonP = make(map[string]int)
	for _, word := range words {
		// try to count how many time 1 word shows up
		if _, ok := nonP[word]; !ok {
			nonP[word] = 1
		} else {
			nonP[word] += 1
		}
	}

	res := 0
	var ok bool = false
	for k, v := range nonP {
		rev := reverse(k)
		if rev == k {
			// if rev = k => this is a palindrome string itself
			// then we have 2 cases:
			// - if there are 2 "ww" => it can be formed into "wwwww"
			// - if there are 3 "ww" => it can only be formed into "wwww" for now because we have not been sure
			// whether we can add another "ww" in the middle or not, we will just turn on the flag that
			// there are a spare "ww" that might can be added later.
			if v%2 == 0 {
				res += v
			} else {
				res += v / 2 * 2
				ok = true
			}
		} else {
			// if it s not a palindrome, check to see if its rev exists and plus the result accordingly
			v1, ok := nonP[rev]
			if !ok {
				continue
			} else {
				res += min(v1, v) * 2
				delete(nonP, k)
				delete(nonP, rev)
			}
		}
	}
	res = res * 2
	if ok {
		res = res + 2
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome([]string{"lc", "cl", "gg"}))
	fmt.Println(longestPalindrome([]string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}))
	fmt.Println(longestPalindrome([]string{"cc", "ll", "xx"}))
	fmt.Println(longestPalindrome([]string{"ab", "ty", "yt", "lc", "cl", "ab"}))
}
