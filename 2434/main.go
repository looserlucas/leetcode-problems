package main

import (
	"fmt"
	"strings"
)

func isSmallest(count []int, c byte) bool {
	j := int(c) - int('a')
	for i := 0; i < j; i++ {
		if count[i] > 0 {
			return false
		}
	}
	return true
}
func robotWithString(s string) string {
	var count = make([]int, 26)
	for i := 0; i < len(s); i++ {
		count[int(s[i])-int('a')]++
	}
	i := 0
	var t []byte
	var res strings.Builder
	for i < len(s) {
		for len(t) > 0 && isSmallest(count, t[len(t)-1]) {
			x := t[len(t)-1]
			res.WriteString(string(x))
			t = t[:len(t)-1]
		}
		t = append(t, s[i])
		count[int(s[i])-int('a')]--
		i++
	}
	for len(t) > 0 {
		x := t[len(t)-1]
		res.WriteString(string(x))
		t = t[:len(t)-1]
	}
	return res.String()
}

func main() {
	fmt.Println(robotWithString("bac"))
	fmt.Println(robotWithString("zza"))
	fmt.Println(robotWithString("bdda"))
}
