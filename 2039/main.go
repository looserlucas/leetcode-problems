package main

import (
	"container/list"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func networkBecomesIdle(edges [][]int, patience []int) int {
	var adj [][]int = make([][]int, len(patience))
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	var dist []int = make([]int, len(patience))
	for i := 1; i < len(dist); i++ {
		dist[i] = math.MaxInt
	}
	var q = list.New()
	q.PushBack(0)
	for q.Len() > 0 {
		top := q.Front()
		u := top.Value.(int)
		for _, e := range adj[u] {
			if dist[e] == math.MaxInt {
				dist[e] = dist[u] + 1
				q.PushBack(e)
			}
		}
		q.Remove(top)
	}
	var res = 0
	for i := 1; i < len(patience); i++ {
		distance := dist[i] * 2
		messageNum := (distance - 1) / patience[i]
		res = max(res, distance+messageNum*patience[i]+1)
	}
	return res
}

func main() {
	//fmt.Println(networkBecomesIdle([][]int{{0, 1}, {1, 2}}, []int{0, 2, 1}))
	//fmt.Println(networkBecomesIdle([][]int{{0, 1}, {0, 2}, {1, 2}}, []int{0, 10, 10}))
}
