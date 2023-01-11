package main

import (
	"fmt"
	"strconv"
	"strings"
)

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	var a []byte = make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		a[i] = s[i]
	}
	end := len(a) - 1
	for i := len(a) - 1; i > 0; i-- {
		if a[i] < a[i-1] {
			end = i - 1
			a[i-1]--
		}
	}
	for i := end + 1; i < len(a); i++ {
		a[i] = '9'
	}
	var res strings.Builder
	for i := 0; i < len(a); i++ {
		res.WriteByte(a[i])
	}
	r, _ := strconv.Atoi(res.String())
	return r
}

func main() {
	fmt.Println("vim-go")
}
