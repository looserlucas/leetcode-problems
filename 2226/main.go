package main

import "fmt"

func check(a []int, k int64, can int) bool {
	sum := int64(0)
	for i := 0; i < len(a); i++ {
		sum += int64(a[i] / can)
	}
	return sum >= k
}

func maximumCandies(a []int, k int64) int {
	max := 0
	total := 0
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
		total += a[i]
	}
	if int64(total) < k {
		return 0
	}
	hi := max
	lo := 1
	for lo < hi {
		mid := (hi + lo) / 2
		if check(a, k, mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if check(a, k, lo) {
		return lo
	} else {
		return lo - 1
	}
}

func main() {
	fmt.Println(maximumCandies([]int{5, 8, 6}, 3))
	fmt.Println(maximumCandies([]int{9, 10, 1, 2, 10, 9, 9, 10, 2, 2}, 3))
	fmt.Println(maximumCandies([]int{2, 5}, 11))
}
