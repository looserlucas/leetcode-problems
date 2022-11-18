package main

import "fmt"

func greedy(s string, r string, x int) (string, int) {
	var des string
	des = des + string(s[0])
	res := 0
	i := 1
	for i < len(s) {
		if len(des) > 0 && s[i] == r[1] && des[len(des)-1] == r[0] {
			des = des[:len(des)-1]
			res += x
		} else {
			des = des + string(s[i])
		}
		i++
	}
	return des, res
}
func maximumGain(s string, x int, y int) int {
	var a, b string = "ab", "ba"
	if x < y {
		a, b = b, a
		x, y = y, x
	}
	newS, res1 := greedy(s, a, x)
	_, res2 := greedy(newS, b, y)
	return res1 + res2
}

func main() {
	fmt.Println(maximumGain("cdbcbbaaabab", 4, 5))
	fmt.Println(maximumGain("aabbaaxybbaabb", 5, 4))
}
