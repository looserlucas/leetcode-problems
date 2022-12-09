package main

import "fmt"

var rank, parent []int

func find(i int) int {
	if parent[i] == i {
		return i
	} else {
		p := find(parent[i])
		parent[i] = p
		return p
	}
}

func union(i, j int) {
	ip := find(i)
	jp := find(j)
	if ip == jp {
		return
	}

	ir := rank[ip]
	jr := rank[jp]
	if ir < jr {
		parent[ip] = jp
	} else {
		parent[jp] = ip
	}
	if ir == jr {
		rank[ir]++
	}
}

func makeConnected(n int, connections [][]int) int {
	// find the redundance cables
	// find how many disconnected cluster
	// disconnect cluster - 1 <= redundance cables
	rank = make([]int, n)
	parent = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	rc := 0
	for _, cable := range connections {
		if find(cable[0]) != find(cable[1]) {
			union(cable[0], cable[1])
		} else {
			rc++
		}
	}

	fmt.Println(rc)
	c := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Println(i, find(i))
		c[find(i)]++
	}

	if len(c)-1 <= rc {
		return len(c) - 1
	} else {
		return -1
	}
}

func main() {
	fmt.Println(makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
}
