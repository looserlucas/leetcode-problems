package main

import (
	"fmt"
	"math"
	"sort"
)

type DisjoinSet struct {
	// parent[i] is the parent of node i
	parent []int
	// rank[i] is the rank of the tree with root i
	rank []int
}

func (d *DisjoinSet) find(i int) int {
	if d.parent[i] == i {
		return i
	} else {
		res := d.find(d.parent[i])
		// compress parent
		d.parent[i] = res
		return res
	}
}

func (d *DisjoinSet) union(i, j int) {
	irep := d.find(i)
	jrep := d.find(j)
	if irep == jrep {
		return
	}

	irank := d.rank[irep]
	jrank := d.rank[jrep]

	// merge smaller rank tree -> bigger rank tree
	// if irank == jrank, merge 2 tree and increase the rank of the new tree by 1
	if irank < jrank {
		d.parent[irep] = jrep
	} else if jrank < irank {
		d.parent[jrep] = irep
	} else {
		d.parent[irep] = jrep
		d.rank[jrep]++
	}
}

func CreateDisjoinSet(length int) *DisjoinSet {
	parent := make([]int, length)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	rank := make([]int, length)
	return &DisjoinSet{
		parent: parent,
		rank:   rank,
	}
}

func kruskal(n int, edges [][]int, blockEdges int, initEdge int) int {
	d := CreateDisjoinSet(n)

	var ans int = 0
	if initEdge != -1 {
		ans += edges[initEdge][2]
		d.union(edges[initEdge][0], edges[initEdge][1])
	}

	for i := 0; i < len(edges); i++ {
		if i == blockEdges {
			continue
		}
		x := edges[i][0]
		y := edges[i][1]
		w := edges[i][2]
		if d.find(x) != d.find(y) {
			d.union(x, y)
			ans += w
		}
	}

	// check if we remove that edge, would it make the graph disconnected
	// if so, return MaxInt
	for i := 0; i < n; i++ {
		if d.find(i) != d.find(0) {
			return math.MaxInt
		}
	}

	return ans
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	for i := 0; i < len(edges); i++ {
		edges[i] = append(edges[i], i)
	}
	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	originalMST := kruskal(n, edges, -1, -1)
	var criticalArr []int
	var psuedoCriticalArr []int

	for i := 0; i < len(edges); i++ {
		if kruskal(n, edges, i, -1) > originalMST {
			criticalArr = append(criticalArr, edges[i][3])
		} else if kruskal(n, edges, -1, i) == originalMST {
			psuedoCriticalArr = append(psuedoCriticalArr, edges[i][3])
		}
	}

	var finalRes [][]int
	finalRes = append(finalRes, criticalArr)
	finalRes = append(finalRes, psuedoCriticalArr)
	return finalRes
}

func main() {
	fmt.Println(findCriticalAndPseudoCriticalEdges(5, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}))
}
