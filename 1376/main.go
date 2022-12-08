package main

import "fmt"

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	var adj = make([][]int, len(manager))
	for i := 0; i < len(manager); i++ {
		if manager[i] != -1 {
			adj[manager[i]] = append(adj[manager[i]], i)
		}
	}
	return dfs(adj, informTime, headID)
}

func dfs(adj [][]int, informTime []int, i int) int {
	var time = 0
	for _, v := range adj[i] {
		t := dfs(adj, informTime, v)
		if t > time {
			time = t
		}
	}
	return time + informTime[i]
}

func main() {
	fmt.Println(numOfMinutes(1, 0, []int{-1}, []int{0}))
	fmt.Println(numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}))
}
