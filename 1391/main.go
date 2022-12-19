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

	ri := rank[ip]
	rj := rank[jp]
	if ri < rj {
		parent[ip] = jp
	} else {
		parent[jp] = ip
		if ri == rj {
			rank[ip]++
		}
	}
}

func rep(i, j, m int) int {
	return i*m + j
}
func hasValidPath(a [][]int) bool {
	n := len(a)
	m := len(a[0])

	rank = make([]int, n*m)
	parent = make([]int, n*m)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch a[i][j] {
			case 1:
				if j-1 >= 0 && (a[i][j-1] == 1 || a[i][j-1] == 4 || a[i][j-1] == 6) {
					union(rep(i, j-1, m), rep(i, j, m))
				}
				if j+1 < m && (a[i][j+1] == 1 || a[i][j+1]%2 == 3 || a[i][j+1] == 5) {
					union(rep(i, j+1, m), rep(i, j, m))
				}
			case 2:
				if i-1 >= 0 && (a[i-1][j] == 2 || a[i-1][j] == 3 || a[i-1][j] == 4) {
					union(rep(i-1, j, m), rep(i, j, m))
				}
				if i+1 < n && (a[i+1][j] == 2 || a[i+1][j] == 5 || a[i+1][j] == 6) {
					union(rep(i+1, j, m), rep(i, j, m))
				}
			case 3:
				if j-1 >= 0 && (a[i][j-1] == 1 || a[i][j-1] == 4 || a[i][j-1] == 6) {
					union(rep(i, j-1, m), rep(i, j, m))
				}
				if i+1 < n && (a[i+1][j] == 2 || a[i+1][j] == 5 || a[i+1][j] == 6) {
					union(rep(i+1, j, m), rep(i, j, m))
				}
			case 4:
				if j+1 < m && (a[i][j+1] == 1 || a[i][j+1] == 3 || a[i][j+1] == 5) {
					union(rep(i, j+1, m), rep(i, j, m))
				}
				if i+1 < n && (a[i+1][j] == 2 || a[i+1][j] == 5 || a[i+1][j] == 6) {
					union(rep(i+1, j, m), rep(i, j, m))
				}
			case 5:
				if i-1 >= 0 && (a[i-1][j] == 2 || a[i-1][j] == 3 || a[i-1][j] == 4) {
					union(rep(i-1, j, m), rep(i, j, m))
				}
				if j-1 >= 0 && (a[i][j-1] == 1 || a[i][j-1] == 4 || a[i][j-1] == 6) {
					union(rep(i, j-1, m), rep(i, j, m))
				}
			case 6:
				if i-1 >= 0 && (a[i-1][j] == 2 || a[i-1][j] == 3 || a[i-1][j] == 4) {
					union(rep(i-1, j, m), rep(i, j, m))
				}
				if j+1 < m && (a[i][j+1] == 1 || a[i][j+1] == 3 || a[i][j+1] == 5) {
					union(rep(i, j+1, m), rep(i, j, m))
				}
			}
		}
	}
	return find(rep(0, 0, m)) == find(rep(n-1, m-1, m))
}

func main() {
	fmt.Println(hasValidPath([][]int{{2, 4, 3}, {6, 5, 2}}))
	fmt.Println(hasValidPath([][]int{{1, 2, 1}, {1, 2, 1}}))
	fmt.Println(hasValidPath([][]int{{1, 1, 2}}))
}
