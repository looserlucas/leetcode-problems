package main

import "fmt"

type DisjointSet struct {
	rank   []int
	parent []int
}

func (s *DisjointSet) find(i int) int {
	if s.parent[i] == i {
		return i
	} else {
		p := s.find(s.parent[i])
		s.parent[i] = p
		return p
	}
}

func (s *DisjointSet) union(i, j int) {
	ip := s.find(i)
	jp := s.find(j)
	if ip == jp {
		return
	}

	ir := s.rank[ip]
	jr := s.rank[jp]
	if ir < jr {
		s.parent[ip] = jp
	} else {
		s.parent[jp] = ip
	}
	if ir == jr {
		s.rank[ip]++
	}
}

func CreateDisjointSet(n int) *DisjointSet {
	rank := make([]int, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &DisjointSet{
		rank:   rank,
		parent: parent,
	}
}

func countServers(a [][]int) int {
	n := len(a)
	m := len(a[0])

	for i := 0; i < n; i++ {
		a[i] = append(a[i], 1)
	}
	var tmp []int
	for j := 0; j <= m; j++ {
		tmp = append(tmp, 1)
	}
	a = append(a, tmp)

	n = len(a)
	m = len(a[0])

	rSet := CreateDisjointSet(n * m)
	cSet := CreateDisjointSet(n * m)

	for i := 0; i < n; i++ {
		last := -1
		count := 0
		for j := m - 1; j >= 0; j-- {
			if a[i][j] == 1 {
				if last == -1 {
					last = i*m + j
				} else {
					rSet.union(last, i*m+j)
				}
				count++
			}
		}
		if count == 2 {
			for j := m - 1; j >= 0; j-- {
				rSet.parent[i*m+j] = i*m + j
				rSet.rank[i*m+j] = 0
			}
		}
	}

	for j := 0; j < m; j++ {
		last := -1
		count := 0
		for i := n - 1; i >= 0; i-- {
			if a[i][j] == 1 {
				if last == -1 {
					last = i*m + j
				} else {
					cSet.union(last, i*m+j)
				}
				count++
			}
		}
		if count == 2 {
			for i := n - 1; i >= 0; i-- {
				cSet.parent[i*m+j] = i*m + j
				cSet.rank[i*m+j] = 0
			}
		}
	}

	var res = 0
	for i := 0; i < n-1; i++ {
		for j := 0; j < m-1; j++ {
			if a[i][j] == 1 && (cSet.find(i*m+j) != i*m+j || rSet.find(i*m+j) != i*m+j) {
				res++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(countServers([][]int{{1, 0}, {0, 1}}))
	fmt.Println(countServers([][]int{{1, 0}, {1, 1}}))
	fmt.Println(countServers([][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}))
}
