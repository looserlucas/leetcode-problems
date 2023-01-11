package main

import (
	"fmt"
	"sort"
)

var MODN = int(1e9 + 7)

func numSubseq(a []int, target int) int {
	n := len(a)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	var pow = make([]int, n)
	pow[0] = 1
	for i := 1; i < n; i++ {
		pow[i] = pow[i-1] * 2 % MODN
	}
	l := 0
	r := n - 1
	var res = 0
	for l < r {
		if a[l]+a[r] > target {
			r -= 1
		} else {
			res = (res + pow[r-l]) % MODN
			l++
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
