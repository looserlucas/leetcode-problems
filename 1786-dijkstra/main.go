package main

import (
	"container/heap"
	"fmt"
	"math"
)

// An IntHeap is a min-heap of ints.
type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func CreateHeap() *IntHeap {
	h := &IntHeap{}
	heap.Init(h)
	return h
}

var dp []int

func countRestrictedPaths(n int, edges [][]int) int {
	var adjV = make(map[int][][]int)
	for i := 0; i < len(edges); i++ {
		v, ok := adjV[edges[i][0]]
		if !ok {
			var tmp [][]int
			tmp = append(tmp, []int{edges[i][1], edges[i][2]})
			adjV[edges[i][0]] = tmp
		} else {
			v = append(v, []int{edges[i][1], edges[i][2]})
			adjV[edges[i][0]] = v
		}

		v, ok = adjV[edges[i][1]]
		if !ok {
			var tmp [][]int
			tmp = append(tmp, []int{edges[i][0], edges[i][2]})
			adjV[edges[i][1]] = tmp
		} else {
			v = append(v, []int{edges[i][0], edges[i][2]})
			adjV[edges[i][1]] = v
		}
	}

	// dijkstra algorithm here
	var dist = make([]int, n+1)
	source := n
	h := CreateHeap()
	for i := 1; i <= n; i++ {
		if i != source {
			dist[i] = math.MaxInt
		} else {
			heap.Push(h, []int{source, 0})
		}
	}

	for h.Len() > 0 {
		u := h.Pop().([]int)[0]
		for _, edge := range adjV[u] {
			v := edge[0]
			w := edge[1]

			if dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
				heap.Push(h, []int{v, dist[v]})
			}
		}
	}

	var visited = make([]bool, n+1)
	dp = make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}

	return dfs(adjV, dist, visited, 1, n)
}

var MODN = int(1e9 + 7)

func dfs(adjV map[int][][]int, dist []int, visited []bool, now, target int) int {
	if dp[now] != -1 {
		return dp[now]
	}
	if now == target {
		return 1
	}
	res := 0
	for _, edge := range adjV[now] {
		v := edge[0]
		if !visited[v] && dist[now] > dist[v] {
			visited[v] = true
			res = (res + dfs(adjV, dist, visited, v, target)) % MODN
			visited[v] = false
		}
	}
	if dp[now] == -1 {
		dp[now] = res
	}
	return res
}

func main() {
	fmt.Println(countRestrictedPaths(5, [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}))
	fmt.Println(countRestrictedPaths(7, [][]int{{1, 3, 1}, {4, 1, 2}, {7, 3, 4}, {2, 5, 3}, {5, 6, 1}, {6, 7, 2}, {7, 5, 3}, {2, 6, 4}}))
}
