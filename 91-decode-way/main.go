package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		if string(s[0]) == "0" {
			return 0
		} else {
			return 1
		}

	}
	var dp = make([]int, len(s)+1)

	if string(s[0]) == "0" {
		dp[0] = 0
	} else {
		dp[0] = 1
	}

	if string(s[0]) == "1" || (string(s[0]) == "2" && string(s[1]) <= "6") {
		if string(s[1]) == "0" {
			dp[1] = 1
		} else {
			dp[1] = 2
		}
	} else {
		if string(s[0]) == "0" || string(s[1]) == "0" {
			return 0
		}
		dp[1] = 1
	}

	for i := 2; i < len(s); i++ {
		if string(s[i]) == "0" {
			if string(s[i-1]) == "1" || string(s[i-1]) == "2" {
				dp[i] = dp[i-2]
			} else {
				return 0
			}
		} else {
			if string(s[i-1]) == "1" || (string(s[i-1]) == "2" && string(s[i]) <= "6") {
				dp[i] = dp[i-2] + dp[i-1]
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
	return dp[len(s)-1]
}

func main() {
	fmt.Println("vim-go")
}
