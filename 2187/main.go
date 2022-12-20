package main

import (
	"fmt"
	"math"
)

func check(a []int, target int, time int64) bool {
	sum := int64(0)
	for i := 0; i < len(a); i++ {
		sum += time / int64(a[i])
	}
	return sum >= int64(target)
}

func minimumTime(a []int, totalTrips int) int64 {
	var minA int = math.MaxInt
	for i := 0; i < len(a); i++ {
		if a[i] < minA {
			minA = a[i]
		}
	}

	var hi, lo int64
	hi = int64(minA) * int64(totalTrips)
	lo = int64(0)
	for lo < hi {
		mid := (lo + hi) / 2
		if check(a, totalTrips, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
	fmt.Println(minimumTime([]int{2}, 1))
}
