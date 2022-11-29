package main

import "fmt"

func strWithout3a3b(a int, b int) string {
	aChar := "a"
	bChar := "b"
	if b > a {
		aChar, bChar = bChar, aChar
		a, b = b, a
	}
	var res string
	for a > 0 {
		res = res + aChar
		a--
		if a > b {
			res = res + aChar
			a--
		}
		if b > 0 {
			res = res + bChar
			b--
		}
	}
	return res
}
func main() {
	fmt.Println("vim-go")
}
