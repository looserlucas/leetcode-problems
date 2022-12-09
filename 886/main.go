package main

import (
	"container/list"
	"fmt"
)

func possibleBipartition(n int, dislikes [][]int) bool {
	var adj [][]int = make([][]int, n+1)
	for _, e := range dislikes {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	var color []int
	for i := 0; i <= n; i++ {
		color = append(color, -1)
	}
	for i := 1; i <= n; i++ {
		if color[i] == -1 {
			q := list.New()
			q.PushBack(i)
			color[i] = 0
			for q.Len() > 0 {
				top := q.Front()
				u := top.Value.(int)
				for _, v := range adj[u] {
					if color[v] == color[u] {
						return false
					}
					if color[v] == -1 {
						color[v] = 1 - color[u]
						q.PushBack(v)
					}
				}
				q.Remove(top)
			}
		}
	}
	return true
}

func main() {
	fmt.Println(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
}
