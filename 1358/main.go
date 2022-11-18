package main

import "fmt"

func numberOfSubstrings(s string) int {
	count := make([]int, 3)
	i := 0
	res := 0
	for j := 0; j < len(s); j++ {
		count[int(s[j])-'a']++
		for i < len(s) && count[0] > 0 && count[1] > 0 && count[2] > 0 {
			count[s[i]-'a']--
			i++
		}
		res += i
	}
	return res
}

func main() {
	fmt.Println(numberOfSubstrings("abcabc"))
}
