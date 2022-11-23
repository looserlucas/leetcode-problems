package main

import "fmt"

func longestPalindrome(s string) string {
	var res string
	var resLen int
	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > resLen {
				res = s[l : r+1]
				resLen = r - l + 1
			}
			l--
			r++
		}
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > resLen {
				res = s[l : r+1]
				resLen = r - l + 1
			}
			l--
			r++
		}
	}

	return res
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("bbd"))
	fmt.Println(longestPalindrome("bbbbbbbbc"))
	fmt.Println(longestPalindrome("bb"))
	fmt.Println(longestPalindrome("bbb"))
}
