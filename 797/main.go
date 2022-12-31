package main

import "fmt"

var result [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	var nowArr = []int{}
	result = nil
	getToTarget(graph, 0, nowArr)
	return result
}

func getToTarget(graph [][]int, now int, nowArr []int) {
	nowArr = append(nowArr, now)
	if now == len(graph)-1 {
		newArr := make([]int, len(nowArr))
		copy(newArr, nowArr)
		result = append(result, newArr)
		nowArr = nowArr[:len(nowArr)-1]
		return
	}

	for _, v := range graph[now] {
		getToTarget(graph, v, nowArr)
	}
	nowArr = nowArr[:len(nowArr)-1]
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{5, 1}, {5, 3, 7, 2}, {7, 4, 6, 3, 5}, {4, 7, 5}, {6, 5, 7}, {7, 6}, {7}, {}}))
}
