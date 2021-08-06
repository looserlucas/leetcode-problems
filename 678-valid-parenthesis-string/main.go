package main

import "fmt"

func checkValidString(s string) bool {
	min := 0
	max := 0
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if c == "(" {
			min++
		} else {
			min--
		}
		if c == ")" {
			max--
		} else {
			max++
		}
		if max < 0 {
			break
		}
		if min < 0 {
			min = 0
		}
	}

	return min == 0
}
func main() {
	s := "(*)("
	fmt.Println(checkValidString(s))
}
