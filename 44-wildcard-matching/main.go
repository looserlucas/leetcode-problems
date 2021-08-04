package main

import "fmt"

func isMatch(s string, p string) bool {
	var memo [][]bool

	var newP string
	var isFirst bool = true
	for i := 0; i < len(p); i++ {
		if string(p[i]) == "*" {
			if isFirst {
				newP = newP + string(p[i])
				isFirst = false
			}
		} else {
			newP = newP + string(p[i])
			isFirst = true
		}
	}
	p = newP
	s = " " + s
	p = " " + p

	for i := 0; i < len(s); i++ {
		var tmp []bool
		for j := 0; j < len(p); j++ {
			tmp = append(tmp, false)
		}
		memo = append(memo, tmp)
	}

	memo[0][0] = true
	if len(p) > 1 && string(p[1]) == "*" {
		memo[0][1] = true
	}

	for i := 1; i < len(s); i++ {
		for j := 1; j < len(p); j++ {
			if s[i] == p[j] || string(p[j]) == "?" {
				memo[i][j] = memo[i-1][j-1]
			} else if string(p[j]) == "*" {
				memo[i][j] = memo[i-1][j] || memo[i][j-1] || memo[i-1][j-1]
			}
		}
	}

	return memo[len(s)-1][len(p)-1]
}

/*
func isMatchMemo(s, p string, memo *[2001][2001]bool, i, j int) {
	if i == len(s) && j == len(p) {
		return memo[i-1][j-1]
	}
}
*/
func main() {
	s := "cb"
	p := "?a"
	fmt.Println(isMatch(s, p))
}
