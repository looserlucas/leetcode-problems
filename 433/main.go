package main

import (
	"fmt"
)

func isNext(g1, g2 string) bool {
	var diff int = 0
	for i := 0; i < len(g1); i++ {
		if g1[i] != g2[i] {
			diff++
		}
	}
	if diff == 1 {
		return true
	} else {
		return false
	}
}

var bank []string
var end string

func minMutation(s string, e string, b []string) int {
	if len(b) == 0 {
		return 0
	}
	if len(b) == 1 {
		return 1
	}

	end = e
	bank = append(bank, s)
	bank = append(bank, b...)
	value := 1 << 0
	res := dfs(0, value)
	if res == 9999 {
		return -1
	} else {
		return res
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func dfs(index int, done int) int {
	if bank[index] == end {
		return 0
	}
	var res int = 9999
	for i := 1; i < len(bank); i++ {
		mask := 1 << i
		if done&mask == 0 && isNext(bank[index], bank[i]) {
			newDone := done | mask
			res = min(res, dfs(i, newDone)+1)
		}
	}
	return res
}
func main() {
	fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))
}
