package main

import "fmt"

func isMatch(s string, p string) bool {
	var memo [][]bool
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
	for i := 2; i < len(p); i++ {
		if string(p[i]) == "*" {
			memo[0][i] = memo[0][i-2]
		}
	}

	for i := 1; i < len(s); i++ {
		for j := 1; j < len(p); j++ {
			if s[i] == p[j] || string(p[j]) == "." {
				memo[i][j] = memo[i-1][j-1]
			} else if string(p[j]) == "*" {
				if j-2 >= 0 {
					memo[i][j] = memo[i][j-2]
				}

				if string(p[j-1]) == "." || p[j-1] == s[i] {
					memo[i][j] = memo[i][j] || memo[i-1][j]
				}
			} else {
				memo[i][j] = false
			}
		}
	}

	return memo[len(s)-1][len(p)-1]
}

func main() {
	s := "mississippi"
	p := "mis*is*p*."
	fmt.Println(isMatch(s, p))
}
