package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	tmp := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			tmp = i
			break
		}
	}

	s = s[tmp:]
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			tmp = i
			break
		}
	}
	s = s[:tmp+1]

	var j = -1
	var i = 0
	for i < len(s) {
		if s[i] == ' ' && j != -1 {
			s = s[:i] + s[i+1:]
			continue
		}
		if s[i] == ' ' && j == -1 {
			j = i
		}

		if s[i] != ' ' {
			j = -1
		}
		i++
	}

	var res string
	arr := strings.Split(s, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		res += arr[i] + " "
	}
	return res[:len(res)-1]
}

func main() {
	fmt.Println(reverseWords("  a good   example  "))
}
