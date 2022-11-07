package main

import (
	"fmt"
	"strconv"
)

func compareVersion(s1 string, s2 string) int {
	i := 0
	j := 0
	for {
		if i == len(s1) && j == len(s2) {
			break
		}
		s1Cur := 0
		for i < len(s1) && s1[i] != '.' {
			n, _ := strconv.Atoi(string(s1[i]))
			s1Cur = s1Cur*10 + n
			i++
		}
		s2Cur := 0
		for j < len(s2) && s2[j] != '.' {
			n, _ := strconv.Atoi(string(s2[j]))
			s2Cur = s2Cur*10 + n
			j++
		}
		if s1Cur > s2Cur {
			return 1
		} else if s1Cur < s2Cur {
			return -1
		}
		if i < len(s1) {
			i++
		}
		if j < len(s2) {
			j++
		}
	}
	return 0
}

func main() {
	fmt.Println(compareVersion("1.01", "1.001.1"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
	fmt.Println(compareVersion("1.1", "1.0.0"))
}
