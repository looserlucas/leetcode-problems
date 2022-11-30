package main

import (
	"sort"
)

// Kruskal algorithm is to find minimum spanning tree, a spanning tree is
// the tree that includes edges that connect all the vertices of the graph
// And the minimum spanning tree is the tree with minimun sum(edges.weight)
// The idea is greedy, pick the edge with the smallest cost posible first
// To do so, we need to sort the edges by their weight first.
// After that, we goes through the list of edges, greedily pick the current
// edge if adding the current edge does not create a circle in the current spanning
// tree that we are making. We can do so using Disjoint Set.
func kruskal(n int, edges [][]int) int {
	d := CreateDisjoinSet(n)
	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	var ans int = 0
	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]
		w := edges[i][2]
		if d.find(x) != d.find(y) {
			d.union(x, y)
			ans += w
		}
	}

	return ans
}
