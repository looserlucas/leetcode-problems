package main

import "fmt"

func canChange(s string, t string) bool {
	i := 0
	j := 0
	n := len(t)
	for i < n || j < n {
		for i < n && s[i] == '_' {
			i++
		}
		for j < n && t[j] == '_' {
			j++
		}
		if i == n || j == n || s[i] != t[j] || (s[i] == 'L' && i < j) || (s[i] == 'R' && i > j) {
			break
		}
		i++
		j++
	}
	return i == n && j == n
}
func main() {
	fmt.Println("vim-go")
}
