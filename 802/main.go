package main

import "fmt"

var visited []bool
var ok []bool

func eventualSafeNodes(graph [][]int) []int {
	visited = make([]bool, len(graph))
	ok = make([]bool, len(graph))
	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			dfs(graph, i)
		}
	}

	var res []int
	for i := 0; i < len(ok); i++ {
		if ok[i] {
			res = append(res, i)
		}
	}
	return res
}

func dfs(graph [][]int, now int) bool {
	if visited[now] {
		return ok[now]
	}
	if len(graph[now]) == 0 {
		ok[now] = true
		return true
	}

	visited[now] = true
	var res = true
	for _, v := range graph[now] {
		res = res && dfs(graph, v)
		fmt.Println(now, v, res)
	}
	ok[now] = res
	return res
}

func main() {
	//fmt.Println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
	//fmt.Println(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}))
	fmt.Println(eventualSafeNodes([][]int{{}, {0, 2, 3, 4}, {3}, {4}, {}}))
}
