package main

import (
	"container/list"
	"fmt"
)

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	var adj [][]int = make([][]int, numCourses)
	var pre [][]bool = make([][]bool, numCourses)
	var inDegree []int = make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		adj[prerequisites[i][0]] = append(adj[prerequisites[i][0]], prerequisites[i][1])
		inDegree[prerequisites[i][1]]++
	}
	var q = list.New()
	for i := 0; i < numCourses; i++ {
		tmp := make([]bool, numCourses)
		pre[i] = tmp
		if inDegree[i] == 0 {
			q.PushBack(i)
		}
	}
	index := numCourses - 1
	var atIndex = make(map[int]int)
	for q.Len() > 0 {
		now := q.Front()
		u := now.Value.(int)
		for _, v := range adj[u] {
			pre[v][u] = true
			for i := 0; i < numCourses; i++ {
				if pre[u][i] {
					pre[v][i] = true
				}
			}
			inDegree[v]--
			if inDegree[v] == 0 {
				q.PushBack(v)
			}
		}
		atIndex[u] = index
		index--
		q.Remove(now)
	}

	var res []bool
	for _, query := range queries {
		res = append(res, pre[query[1]][query[0]])
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
