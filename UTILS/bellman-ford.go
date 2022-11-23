package main

import (
	"fmt"
	"math"
)

// B take 2 inputs: n - number of vertices, edges are list of edges.
// then implement Bellman-Ford algorithm on the input graph.
// Bellman-Ford is a DP algorithm. dist[i][k] is the minimum path from
// source -> vertice i with at most k edges. But we can disgard k and use just dist[i] instead.
// the shortest path should be include of at most n-1 edges. So this we calculate dist[i] k-1 time
func BellmanFord(n int, edges [][]int, src int) int {
	var dist = make([]int, n+1)
	for i := 1; i <= n-1; i++ {
		dist[i] = math.MaxInt
	}
	dist[src] = 0
	for i := 1; i <= n-1; i++ {
		for j := 0; j < len(edges); j++ {
			u := edges[j][0]
			v := edges[j][1]
			w := edges[j][2]
			if dist[u] != math.MaxInt && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
			}
		}
	}

	for i := 1; i <= n-1; i++ {
		for j := 0; j < len(edges); j++ {
			u := edges[j][0]
			v := edges[j][1]
			w := edges[j][2]
			if dist[u] != math.MaxInt && dist[u]+w < dist[v] {
				fmt.Println("Graph contains negative weight cycle")
				return -1
			}
		}
	}
	fmt.Println(dist)
	return 0
}
