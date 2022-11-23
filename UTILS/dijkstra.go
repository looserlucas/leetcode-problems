package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Dijkstra take 2 inputs: n - number of vertices, edges are list of edges. Convert these 2 into adjecent vertices list (adjV)
// then implement dijkstra algorithm on the input graph (with heap included, heap is used to store vertice + its current shortest path)
func Dijkstra(n int, edges [][]int) int {
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

	fmt.Println(dist)
	return 0
}
