package main

import (
	"fmt"
	"strconv"
)

func maxDiff(num int) int {
	s := strconv.Itoa(num)
	var minI, maxI int = -1, -1
	var minChar, maxChar string
	var minRep, maxRep byte
	var s1, s2 string
	for i := 0; i < len(s); i++ {
		if s[i] != '1' && s[i] != '0' {
			minI = i
			break
		}
	}
	var noReplaceMin, noReplaceMax bool
	if minI == -1 {
		noReplaceMin = true
	} else {
		if minI != 0 {
			minChar = "0"
			minRep = s[minI]
		} else {
			minChar = "1"
			minRep = s[minI]
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] != '9' {
			maxI = i
			break
		}
	}
	if maxI == -1 {
		noReplaceMax = true
	} else {
		maxChar = "9"
		maxRep = s[maxI]
	}

	for i := 0; i < len(s); i++ {
		if !noReplaceMin && s[i] == minRep {
			s1 += minChar
		} else {
			s1 += string(s[i])
		}

		if !noReplaceMax && s[i] == maxRep {
			s2 += maxChar
		} else {
			s2 += string(s[i])
		}
	}
	a, _ := strconv.Atoi(s1)
	b, _ := strconv.Atoi(s2)
	fmt.Println(a, b)
	return b - a
}
func main() {
	fmt.Println(maxDiff(555))
	fmt.Println(maxDiff(9))
	fmt.Println(maxDiff(123456))
	fmt.Println(maxDiff(999999))
	fmt.Println(maxDiff(111111))
	fmt.Println(maxDiff(999189))
}
