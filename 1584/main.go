package main

import "sort"

type DisjoinSet struct {
	rank   []int
	parent []int
}

func (s *DisjoinSet) find(i int) int {
	if s.parent[i] == i {
		return i
	} else {
		res := s.find(s.parent[i])
		s.parent[i] = res
		return res
	}
}

func (s *DisjoinSet) union(i, j int) {
	iparent := s.find(i)
	jparent := s.find(j)
	if iparent == jparent {
		return
	}

	irank := s.rank[iparent]
	jrank := s.rank[jparent]
	if irank <= jrank {
		s.parent[iparent] = jparent
	} else {
		s.parent[jparent] = iparent
	}
	if irank == jrank {
		s.rank[irank] += 1
	}
}

func CreateDisjoinSet(n int) *DisjoinSet {
	rank := make([]int, n)
	var parent []int
	for i := 0; i < n; i++ {
		parent = append(parent, i)
	}
	return &DisjoinSet{
		rank:   rank,
		parent: parent,
	}
}

func kruskal(n int, edges [][]int) int {
	s := CreateDisjoinSet(n)
	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	var ans int = 0
	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]
		w := edges[i][2]

		if s.find(x) != s.find(y) {
			s.union(x, y)
			ans += w
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func distance(x []int, y []int) int {
	return abs(x[0]-y[0]) + abs(x[1]-y[1])
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	var edges [][]int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, []int{i, j, distance(points[i], points[j])})
		}
	}
	// O(ElogE + ElogV) = O(ElogE) because E is V^2 => logE = logV
	res := kruskal(n, edges)
	return res
}
