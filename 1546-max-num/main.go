package main

import "fmt"

func maxNonOverlapping(n []int, t int) int {
	if len(n) == 0 {
		return 0
	}
	var ps = make([]int, len(n))
	ps[0] = n[0]
	for i := 1; i < len(n); i++ {
		ps[i] = ps[i-1] + n[i]
	}

	var res int
	var i = 0
	var last int = 0
	for i < len(n) {
		for j := i; j >= last; j-- {
			if ps[i]-ps[j]+n[j] == t {
				res++
				fmt.Println(j, i)
				last = i + 1
				break
			}
		}
		i++
	}
	return res
}
func main() {
	fmt.Println(maxNonOverlapping([]int{-1, 3, 5, 1, 4, 2, -9}, 6))
	fmt.Println(maxNonOverlapping([]int{0, 0, 0}, 0))
}
