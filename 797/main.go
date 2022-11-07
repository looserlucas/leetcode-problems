package main

import "fmt"

var result [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	var nowArr = []int{0}
	result = nil
	getToTarget(graph, 0, nowArr)
	return result
}

func getToTarget(graph [][]int, now int, nowArr []int) {
	if now == len(graph)-1 {
		result = append(result, nowArr)
		return
	}

	for _, v := range graph[now] {
		var newArr = make([]int, len(nowArr))
		copy(newArr, nowArr)
		newArr = append(newArr, v)
		getToTarget(graph, v, newArr)
	}
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{5, 1}, {5, 3, 7, 2}, {7, 4, 6, 3, 5}, {4, 7, 5}, {6, 5, 7}, {7, 6}, {7}, {}}))
}
