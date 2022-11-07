package main

import "fmt"

func longestPalindrome(s string) string {
	var dp = make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == s[i] {
			dp[i] = dp[i-1] + 2
		} else if dp[i-1] == 1 && s[i] == s[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}
	fmt.Println(dp)

	return ""
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("bbd"))
	fmt.Println(longestPalindrome("bbbbbbbbc"))
	fmt.Println(longestPalindrome("bb"))
	fmt.Println(longestPalindrome("bbb"))
}
