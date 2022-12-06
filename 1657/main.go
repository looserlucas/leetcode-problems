package main

import (
	"fmt"
)

func closeStrings(word1 string, word2 string) bool {
	c1 := make(map[byte]int)
	c2 := make(map[byte]int)
	if len(word1) != len(word2) {
		return false
	}
	n := len(word1)
	for i := 0; i < n; i++ {
		c1[word1[i]]++
		c2[word2[i]]++
	}
	if len(c1) != len(c2) {
		return false
	}
	freq1 := make(map[int]int)
	freq2 := make(map[int]int)
	for k, v := range c1 {
		if _, ok := c2[k]; !ok {
			return false
		}
		freq1[v]++
	}
	for k, v := range c2 {
		if _, ok := c1[k]; !ok {
			return false
		}
		freq2[v]++
	}

	for k, _ := range freq1 {
		if freq1[k] != freq2[k] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(closeStrings("abc", "bca"))
	fmt.Println(closeStrings("a", "aa"))
	fmt.Println(closeStrings("cabbba", "abbccc"))
	fmt.Println(closeStrings("cabbbb", "abbccc"))
}
