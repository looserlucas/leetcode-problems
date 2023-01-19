package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func minFlipsMonoIncr(s string) int {
	total0 := 0
	total1 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			total1++
		} else {
			total0++
		}
	}
	var res int = total0
	c1 := 0
	c0 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			c0++
		} else {
			c1++
		}
		res = min(res, total0-c0+c1)
	}
	return res
}

func main() {
	/*
		fmt.Println(minFlipsMonoIncr("00110"))
		fmt.Println(minFlipsMonoIncr("010110"))
		fmt.Println(minFlipsMonoIncr("00011000"))
		fmt.Println(minFlipsMonoIncr("11011"))
	*/
	fmt.Println(minFlipsMonoIncr("1111001110"))
}
