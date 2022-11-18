package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxScore(cardPoints []int, k int) int {
	a := append(cardPoints, cardPoints...)
	i := len(cardPoints) - k
	j := len(cardPoints)
	s := 0
	for k := i; k < j; k++ {
		s += a[k]
	}
	res := 0
	res = max(res, s)
	for j < len(cardPoints)+k {
		fmt.Println(res)
		s += a[j]
		s -= a[i]
		res = max(res, s)
		i++
		j++
	}
	return res
}
func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
}
