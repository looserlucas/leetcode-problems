package main

import (
	"container/list"
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		inDegree[prerequisites[i][0]]++
		adj[prerequisites[i][1]] = append(adj[prerequisites[i][1]], prerequisites[i][0])
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
	if len(res) < numCourses {
		return nil
	} else {
		return res
	}
}

func main() {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
