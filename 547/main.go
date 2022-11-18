package main

import "fmt"

var done []bool

func findCircleNum(isConnected [][]int) int {
	done = nil
	for i := 0; i < len(isConnected); i++ {
		done = append(done, false)
	}
	res := 0
	for i := 0; i < len(isConnected); i++ {
		if !done[i] {
			traverse(i, isConnected)
			res++
		}
	}
	return res
}

func traverse(index int, isConnected [][]int) {
	done[index] = true
	for i := 0; i < len(isConnected[index]); i++ {
		if !done[i] && isConnected[index][i] == 1 {
			traverse(i, isConnected)
		}
	}
}

func main() {
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}
