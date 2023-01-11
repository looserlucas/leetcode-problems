package main

import (
	"fmt"
	"strings"
)

func secondsToRemoveOccurrences(s string) int {
	res := 0
	changed := true
	for changed {
		changed = false
		var newS strings.Builder
		for i := 0; i < len(s)-1; i++ {
			if s[i] == '0' && s[i+1] == '1' {
				newS.WriteByte('1')
				newS.WriteByte('0')
				i++
				changed = true
			} else {
				newS.WriteByte(s[i])
			}
		}
		if newS.Len() < len(s) {
			newS.WriteByte(s[len(s)-1])
		}
		if changed {
			res++
		}
		s = newS.String()
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
