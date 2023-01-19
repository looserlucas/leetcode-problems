package main

import "fmt"

func subarraysDivByK(a []int, k int) int {
	var pre = make([]int, len(a))
	pre[0] = a[0]
	for i := 1; i < len(a); i++ {
		pre[i] = pre[i-1] + a[i]
	}
	var res = 0
	var seenMod = make(map[int]int)
	for i := 0; i < len(a); i++ {
		m := pre[i] % k
		if m == 0 {
			seenMod[m]++
			res += seenMod[m]
		} else if m > 0 {
			if seenMod[m] > 0 {
				res += seenMod[m]
			}
			if seenMod[-(k-m)] > 0 {
				res += seenMod[-(k - m)]
			}
			seenMod[m]++
		} else {
			if seenMod[m] > 0 {
				res += seenMod[m]
			}
			if seenMod[k+m] > 0 {
				res += seenMod[k+m]
			}
			seenMod[m]++
		}
	}
	return res
}

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}
