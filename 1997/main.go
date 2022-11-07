package main

import "fmt"

const MODN = 1e9 + 7

// first[i] is the first time we visit i
// second[i] is the second time we visit i
// We CAN ONLY move forwart (from i-1 to i) when we visit i-1 the second (even) time.
// first[i] = second[i-1] + 1
// To calculate second[i], we have:
// second[i] = first[i] + d
// d is the time we can loop back from nextVisit[i] -> i, we already calculate this one
// d = first[i] - first[nextVisit[i]]
func firstDayBeenInAllRooms(nextVisit []int) int {
	n := len(nextVisit)
	first := make([]int, n)
	second := make([]int, n)
	first[0] = 0
	second[0] = 1
	for i := 1; i < n; i++ {
		first[i] = (second[i-1] + 1) % MODN
		second[i] = (first[i] + first[i] - first[nextVisit[i]] + 1) % MODN
	}
	return first[n-1] % MODN
}

func main() {
	fmt.Println(firstDayBeenInAllRooms([]int{0, 0}))
	fmt.Println(firstDayBeenInAllRooms([]int{0, 0, 2}))
	fmt.Println(firstDayBeenInAllRooms([]int{0, 1, 2, 0}))
}
