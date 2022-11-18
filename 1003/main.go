package main

import "fmt"

func isValid(s string) bool {
	if len(s) <= 2 {
		return false
	}
	var s1 string
	s1 = s1 + string(s[0]) + string(s[1])
	i := 2
	for i < len(s) {
		if len(s1) >= 2 && s[i] == 'c' && s1[len(s1)-1] == 'b' && s1[len(s1)-2] == 'a' {
			s1 = s1[:len(s1)-2]
		} else {
			s1 += string(s[i])
		}
		fmt.Println(s1)
		i++
	}
	fmt.Println(s1)
	if len(s1) > 0 {
		return false
	} else {
		return true
	}
}
func main() {
	fmt.Println(isValid("aabcbc"))
}
