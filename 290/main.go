package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	wordMap := make(map[byte]string)
	byteMap := make(map[string]byte)
	a := strings.Split(s, " ")
	if len(a) != len(pattern) {
		return false
	}
	for i := 0; i < len(a); i++ {
		_, ok := wordMap[pattern[i]]
		if !ok {
			wordMap[pattern[i]] = a[i]
		}
		_, ok1 := byteMap[a[i]]
		if !ok1 {
			byteMap[a[i]] = pattern[i]
		}
		if byteMap[a[i]] != pattern[i] || wordMap[pattern[i]] != a[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
}
