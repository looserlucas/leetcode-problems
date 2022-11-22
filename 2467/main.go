package main

import (
	"fmt"
	"math"
)

var done []int
var a [][]int
var bobRoute []int
var res int

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	a = make([][]int, len(edges)+1)
	res = -math.MaxInt
	done = make([]int, len(a))
	bobRoute = nil

	for i := 0; i < len(edges); i++ {
		a[edges[i][0]] = append(a[edges[i][0]], edges[i][1])
		a[edges[i][1]] = append(a[edges[i][1]], edges[i][0])
	}

	var now []int
	dfsBob(bob, 0, now)
	fmt.Println(bobRoute)

	var bobDone = make(map[int]int)
	for i := 0; i < len(bobRoute); i++ {
		bobDone[bobRoute[i]] = i
	}

	done = make([]int, len(a))
	dfsAlice(0, 0, 0, bobDone, amount)
	return res
}

func dfsBob(s, t int, now []int) {
	if bobRoute != nil {
		return
	}
	done[s] = 1
	new := append(now, s)
	if s == t {
		bobRoute = new
		return
	}

	for i := 0; i < len(a[s]); i++ {
		if done[a[s][i]] == 0 {
			dfsBob(a[s][i], t, new)
		}
	}
}

func dfsAlice(na int, level int, sum int, bobDone map[int]int, amount []int) {
	done[na] = 1
	v, ok := bobDone[na]
	if !ok {
		sum += amount[na]
	}
	// split cost by half
	if v == level {
		sum += (amount[na] / 2)
	} else if v > level {
		sum += amount[na]
	}

	if len(a[na]) == 1 && done[a[na][0]] == 1 {
		if sum > res {
			res = sum
		}
		return
	}

	for i := 0; i < len(a[na]); i++ {
		if done[a[na][i]] == 0 {
			dfsAlice(a[na][i], level+1, sum, bobDone, amount)
		}
	}
}

func main() {
	//fmt.Println(mostProfitablePath([][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, 3, []int{-2, 4, 2, -4, 6}))
	fmt.Println(mostProfitablePath([][]int{{0, 1}}, 1, []int{-7280, 2350}))
	fmt.Println(mostProfitablePath([][]int{{0, 2}, {1, 4}, {1, 6}, {2, 3}, {2, 8}, {3, 7}, {4, 5}, {6, 7}}, 1, []int{-1410, -9440, 5536, -774, 3044, 7924, -2122, -1192, 9166}))
}
