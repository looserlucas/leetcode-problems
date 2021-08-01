package main

import (
	"fmt"
	"strconv"
)

func splitString(s string) bool {
	return dfs(s, -1)
}

func dfs(s string, prev int) bool {
	cur := 0
	for i := 0; i < len(s); i++ {
		n, _ := strconv.Atoi(string(s[i]))
		cur = cur*10 + n

		// take care of overflow
		if int64(cur) >= 1e10 {
			return false
		}
		if prev == -1 {
			if dfs(s[i+1:], cur) {
				return true
			}
		} else if cur == prev-1 && (dfs(s[i+1:], cur) || i == len(s)-1) {
			return true
		}
	}
	return false
}
func main() {
	s := "9080701"
	fmt.Println(splitString(s))
}
