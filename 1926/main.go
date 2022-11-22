package main

import "fmt"

func nearestExit(maze [][]byte, entrance []int) int {
	n := len(maze)
	m := len(maze[0])
	var done = make([][]bool, n)
	for i := 0; i < len(done); i++ {
		tmp := make([]bool, m)
		done[i] = tmp
	}

	var q [][]int
	q = append(q, []int{entrance[0], entrance[1], 0})
	i := 0
	for i < len(q) {
		c := q[i]
		fmt.Println(q, q[i])
		done[c[0]][c[1]] = true
		if c[2] != 0 && (c[0] == 0 || c[1] == 0 || c[0] == n-1 || c[1] == m-1) {
			return c[2]
		}
		if c[0]-1 >= 0 && maze[c[0]-1][c[1]] != '+' && !done[c[0]-1][c[1]] {
			q = append(q, []int{c[0] - 1, c[1], c[2] + 1})
		}
		if c[0]+1 < n && maze[c[0]+1][c[1]] != '+' && !done[c[0]+1][c[1]] {
			q = append(q, []int{c[0] + 1, c[1], c[2] + 1})
		}
		if c[1]-1 >= 0 && maze[c[0]][c[1]-1] != '+' && !done[c[0]][c[1]-1] {
			q = append(q, []int{c[0], c[1] - 1, c[2] + 1})
		}
		if c[1]+1 < m && maze[c[0]][c[1]+1] != '+' && !done[c[0]][c[1]+1] {
			q = append(q, []int{c[0], c[1] + 1, c[2] + 1})
		}
		i++
	}
	return -1
}
func main() {
	//fmt.Println(nearestExit([][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}, []int{1, 2}))
	fmt.Println(nearestExit([][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}, []int{1, 0}))
}
