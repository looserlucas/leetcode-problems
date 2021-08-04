package main

import "fmt"

const moder = int(1e9 + 7)

func countPairs(deliciousness []int) int {
	dd := make(map[int]int)
	var res int
	for i := 0; i < len(deliciousness); i++ {
		power := 1
		for j := 0; j < 22; j++ {
			if _, ok := dd[power-deliciousness[i]]; ok {
				res += dd[power-deliciousness[i]]
				res %= moder
			}
			power *= 2
		}
		if _, ok := dd[deliciousness[i]]; ok {
			dd[deliciousness[i]]++
		} else {
			dd[deliciousness[i]] = 1
		}
	}
	return res
}
func main() {
	deliciousness := []int{1, 1, 1, 3, 3, 3, 7}
	fmt.Println(countPairs(deliciousness))
}
