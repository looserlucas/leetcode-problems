package main

import "fmt"

func checkBouquet(bloomday []int, now int, k int, m int) bool {
	count := 0
	finalCount := 0
	for i := 0; i < len(bloomday); i++ {
		if bloomday[i] <= now {
			count++
			if count == k {
				finalCount++
				count = 0
			}
		} else {
			count = 0
		}
	}
	//fmt.Println(now, finalCount, m)

	if finalCount < m {
		return true
	} else {
		return false
	}
}

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}

	maxValue := -1
	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] > maxValue {
			maxValue = bloomDay[i]
		}
	}

	l := 1
	r := maxValue
	for l < r {
		mid := (l + r) / 2
		if checkBouquet(bloomDay, mid, k, m) {
			l = mid + 1
		} else {
			r = mid
		}
		fmt.Println(l, r)
	}
	return l
}

func main() {
	fmt.Println(minDays([]int{1, 10, 3, 10, 2}, 3, 1))
	//fmt.Println(minDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3))
}
