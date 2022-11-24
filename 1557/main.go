package main

import "fmt"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	var inDegree = make([]int, n)
	for i := 0; i < len(edges); i++ {
		inDegree[edges[i][1]]++
	}
	var res []int
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
