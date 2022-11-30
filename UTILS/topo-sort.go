package main

import "container/list"

func topoSortQueue(n int, edges [][]int) []int {
	inDegree := make([]int, n)
	adj := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		inDegree[edges[i][1]]++
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][0])
	}
	q := list.New()
	for i := 0; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			q.PushBack(i)
		}
	}

	var res []int
	for q.Len() > 0 {
		now := q.Front().Value.(int)
		res = append(res, now)
		for i := 0; i < len(adj[now]); i++ {
			inDegree[adj[now][i]]--
			if inDegree[adj[now][i]] == 0 {
				q.PushBack(adj[now][i])
			}
		}
		q.Remove(q.Front())
	}
	if len(res) < n {
		return nil
	} else {
		return res
	}
}

var flag bool

func topoDFS(n int, edges [][]int) bool {
	flag = false
	adj := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
	}
	var visit = make([]int, n)
	var q []int = nil
	for i := 0; i < n; i++ {
		if visit[i] == 0 {
			dfsTopo(adj, i, visit, &q)
			if flag {
				return false
			}
		}
	}
	return len(q) == n
}

func dfsTopo(adj [][]int, i int, visit []int, q *[]int) {
	visit[i] = 1
	for j := 0; j < len(adj[i]); j++ {
		if visit[adj[i][j]] == 1 {
			flag = true
			return
		}
		if visit[adj[i][j]] == 0 {
			dfsTopo(adj, adj[i][j], visit, q)
		}
	}
	(*q) = append(*q, i)
	visit[i] = 2
}
