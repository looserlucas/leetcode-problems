package main

import "fmt"

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	l := 0
	r := int(1e9)
	for l < r {
		m := (l + r) / 2
		if helper(divisor1, divisor2, uniqueCnt1, uniqueCnt2, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func helper(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int, m int) bool {
	d1 := m / divisor1
	d2 := m / divisor2
	d3 := m / (divisor1 * divisor2)
	if d1-d3 >= uniqueCnt1 && d2-d3 >= uniqueCnt2 {
		return true
	}
	c1 := 0
	if uniqueCnt1-(d1-d3) > 0 {
		c1 = uniqueCnt1 - (d1 - d3)
	}
	c2 := 0
	if uniqueCnt2-(d2-d3) > 0 {
		c2 = uniqueCnt2 - (d2 - d3)
	}
	if c1+c2 <= d3 {
		return true
	}
	return false
}

func main() {
	fmt.Println("vim-go")
}
