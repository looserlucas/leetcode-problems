package main

import "fmt"

var id = 0
var ids []int
var low []int

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func dfs(adjV map[int][]int, i int, parent int) {
	ids[i] = id
	low[i] = id
	id++

	for _, v := range adjV[i] {
		if ids[v] == -1 {
			dfs(adjV, v, i)
		}
		if parent != -1 {
			if v != parent {
				low[i] = min(low[i], low[v])
			}
		} else {
			low[i] = min(low[i], low[v])
		}
	}
}

func criticalConnections(n int, edges [][]int) [][]int {
	var adjV = make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		v, ok := adjV[edges[i][0]]
		if !ok {
			var tmp []int
			tmp = append(tmp, edges[i][1])
			adjV[edges[i][0]] = tmp
		} else {
			v = append(v, edges[i][1])
			adjV[edges[i][0]] = v
		}

		v, ok = adjV[edges[i][1]]
		if !ok {
			var tmp []int
			tmp = append(tmp, edges[i][0])
			adjV[edges[i][1]] = tmp
		} else {
			v = append(v, edges[i][0])
			adjV[edges[i][1]] = v
		}
	}

	ids = make([]int, n)
	low = make([]int, n)
	for i := 0; i < len(ids); i++ {
		ids[i] = -1
	}

	for i := 0; i < n; i++ {
		if ids[i] == -1 {
			dfs(adjV, i, -1)
		}
	}

	var res [][]int
	for i := 0; i < len(edges); i++ {
		if low[edges[i][0]] > ids[edges[i][1]] || low[edges[i][1]] > ids[edges[i][0]] {
			res = append(res, edges[i])
		}
	}
	return res
}

func main() {
	fmt.Println(criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
}
