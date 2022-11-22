package main

import "fmt"

func minimumLength(s string) int {
	i := 0
	j := len(s) - 1
	var res = 0
	for i <= j {
		if s[i] != s[j] {
			res = j - i + 1
			break
		}
		if (i == j) && (i+1 == len(s) || (i+1 < len(s) && s[i] != s[i+1])) {
			res = 1
		}
		for j > i && s[j] == s[j-1] {
			j--
		}
		for i < j && s[i] == s[i+1] {
			i++
		}
		i++
		j--
	}

	return res
}

func main() {
	fmt.Println(minimumLength("ca"))
	fmt.Println(minimumLength("cabaabac"))
	fmt.Println(minimumLength("aa"))
	fmt.Println(minimumLength("aabccabba"))
	fmt.Println(minimumLength("bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb"))
}
