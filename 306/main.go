package main

import (
	"fmt"
	"strconv"
)

func isAdditiveNumber(num string) bool {
	return dfs(num, -1, -1)
}

func dfs(s string, prev1, prev2 int) bool {
	cur := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			cur = cur * 10
			continue
		} else {
			n, _ := strconv.Atoi(string(s[i]))
			cur = cur*10 + n
		}
		// take care of overflow
		if int64(cur) >= 1e10 {
			return false
		}

		if prev1 == -1 {
			if dfs(s[i+1:], cur, -1) {
				return true
			}
		} else if prev2 == -1 {
			if dfs(s[i+1:], prev1, cur) {
				return true
			}
		} else if cur == prev1+prev2 && (dfs(s[i+1:], prev2, cur) || i == len(s)-1) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isAdditiveNumber("112358"))
	fmt.Println(isAdditiveNumber("199100199"))
	fmt.Println(isAdditiveNumber("1023"))
	fmt.Println(isAdditiveNumber("101"))
	fmt.Println(isAdditiveNumber("125"))
}
