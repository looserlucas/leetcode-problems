package main

import "fmt"

func longestSubsequence(s string, k int) int {
	val := 0
	cnt := 0
	pow := 1
	for i := len(s) - 1; i >= 0 && val+pow <= k; i-- {
		if s[i] == '1' {
			cnt++
			val += pow
		}
		pow <<= 1
	}
	cntZ := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			cntZ++
		}
	}
	return cnt + cntZ
}
func main() {
	fmt.Println("vim-go")
}
