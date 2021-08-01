package main

import "fmt"

const (
	MOUSE = 1
	CAT   = 2
	DRAW  = 0
)

type MinimaxNode struct {
	c      int
	m      int
	t      int
	winner int
}

func getParent(graph [][]int, m, c, t int) []*MinimaxNode {
	var ans []*MinimaxNode
	// if current turn is cat, then the parent turn is mouse
	if t == 2 {

		for i := 0; i < len(graph[m]); i++ {
			ans = append(ans, &MinimaxNode{
				m: graph[m][i],
				c: c,
				t: 3 - t,
			})
		}
	} else {
		for i := 0; i < len(graph[c]); i++ {
			if graph[c][i] > 0 {
				ans = append(ans, &MinimaxNode{
					m: m,
					c: graph[c][i],
					t: 3 - t,
				})
			}
		}
	}
	return ans
}

func catMouseGame(graph [][]int) int {
	n := len(graph)
	var color [50][50][3]int
	var degree [50][50][3]int

	for m := 0; m < n; m++ {
		for c := 0; c < n; c++ {
			degree[m][c][1] = len(graph[m])
			degree[m][c][2] = len(graph[c])
			for i := 0; i < len(graph[c]); i++ {
				if graph[c][i] == 0 {
					degree[m][c][2]--
					break
				}
			}
		}
	}

	var list []*MinimaxNode
	// add core winning state to
	for i := 0; i < n; i++ {
		for j := 1; j <= 2; j++ {
			color[0][i][j] = MOUSE
			list = append(list, &MinimaxNode{
				c:      i,
				m:      0,
				t:      j,
				winner: MOUSE,
			})
			if i > 0 {
				color[i][i][j] = CAT
				list = append(list, &MinimaxNode{
					c:      i,
					m:      i,
					t:      j,
					winner: CAT,
				})
			}
		}
	}

	d := 0
	for d < len(list) {
		i := list[d].m
		j := list[d].c
		t := list[d].t
		c := list[d].winner
		parents := getParent(graph, i, j, t)
		for k := 0; k < len(parents); k++ {
			i1 := parents[k].m
			j1 := parents[k].c
			t1 := parents[k].t

			if color[i1][j1][t1] == DRAW {
				if t1 == c {
					// if the parent can make the winnin move, then make it
					color[i1][j1][t1] = c
					list = append(list, &MinimaxNode{
						m:      i1,
						c:      j1,
						t:      t1,
						winner: c,
					})
				} else {
					// if not then, mark this child is useless
					degree[i1][j1][t1]--
					// if all children are useless, then this one move will let to the guarantee lost
					if degree[i1][j1][t1] == 0 {
						color[i1][j1][t1] = 3 - t1
						list = append(list, &MinimaxNode{
							m:      i1,
							c:      j1,
							t:      t1,
							winner: 3 - t1,
						})
					}
				}
			}
		}
		d++
	}
	return color[1][2][1]
}

func main() {
	var graph = [][]int{
		{2, 5},
		{3},
		{0, 4, 5},
		{1, 4, 5},
		{2, 3},
		{0, 2, 3},
	}
	fmt.Println(catMouseGame(graph))
}
