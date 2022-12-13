package main

import "fmt"

var rank, parent []int
var visited []bool

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
		rank[ip]++
	}
}

func rep(i, j, m int) int {
	return i*m + j
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	n := len(grid1)
	m := len(grid1[0])
	rank = make([]int, n*m)
	parent = make([]int, n*m)
	visited = make([]bool, n*m)
	for i := 0; i < n*m; i++ {
		parent[i] = i
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid1[i][j] == 1 {
				r := rep(i, j, m)
				if !visited[r] {
					dfs1(grid1, i, j, r)
				}
			}
		}
	}

	visited = make([]bool, n*m)
	var res = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid2[i][j] == 1 {
				r := rep(i, j, m)
				if grid1[i][j] == 1 {
					if !visited[r] {
						p := find(r)
						if dfs2(grid2, i, j, p) {
							res++
						}
					}
				} else {
					if !visited[r] {
						dfs2(grid2, i, j, -1)
					}
				}
			}

		}
	}
	return res
}

func dfs1(a [][]int, i, j, p int) {
	n := len(a)
	m := len(a[0])
	visited[i*m+j] = true
	union(i*m+j, p)
	if i-1 >= 0 && !visited[(i-1)*m+j] && a[i-1][j] == 1 {
		dfs1(a, i-1, j, p)
	}
	if i+1 < n && !visited[(i+1)*m+j] && a[i+1][j] == 1 {
		dfs1(a, i+1, j, p)
	}
	if j+1 < m && !visited[i*m+j+1] && a[i][j+1] == 1 {
		dfs1(a, i, j+1, p)
	}
	if j-1 >= 0 && !visited[i*m+j-1] && a[i][j-1] == 1 {
		dfs1(a, i, j-1, p)
	}
}

func dfs2(a [][]int, i, j, p int) bool {
	n := len(a)
	m := len(a[0])
	visited[i*m+j] = true
	var ok = true
	if p != -1 && find(i*m+j) != p {
		ok = false
	}
	var ok1, ok2, ok3, ok4 bool = true, true, true, true
	if i-1 >= 0 && !visited[(i-1)*m+j] && a[i-1][j] == 1 {
		ok1 = dfs2(a, i-1, j, p)
	}
	if i+1 < n && !visited[(i+1)*m+j] && a[i+1][j] == 1 {
		ok2 = dfs2(a, i+1, j, p)
	}
	if j+1 < m && !visited[i*m+j+1] && a[i][j+1] == 1 {
		ok3 = dfs2(a, i, j+1, p)
	}
	if j-1 >= 0 && !visited[i*m+j-1] && a[i][j-1] == 1 {
		ok4 = dfs2(a, i, j-1, p)
	}

	return ok && ok1 && ok2 && ok3 && ok4
}

func main() {
	fmt.Println(countSubIslands([][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}},
		[][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}))
	fmt.Println(countSubIslands([][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}},
		[][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}}))
	fmt.Println(countSubIslands([][]int{{1, 1, 1, 1, 0, 0}, {1, 1, 0, 1, 0, 0}, {1, 0, 0, 1, 1, 1}, {1, 1, 1, 0, 0, 1}, {1, 1, 1, 1, 1, 0}, {1, 0, 1, 0, 1, 0}, {0, 1, 1, 1, 0, 1}, {1, 0, 0, 0, 1, 1}, {1, 0, 0, 0, 1, 0}, {1, 1, 1, 1, 1, 0}},
		[][]int{{1, 1, 1, 1, 0, 1}, {0, 0, 1, 0, 1, 0}, {1, 1, 1, 1, 1, 1}, {0, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 1, 0}, {0, 1, 1, 1, 1, 1}, {1, 1, 0, 1, 1, 1}, {1, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1}, {1, 0, 0, 1, 0, 0}}))

}
