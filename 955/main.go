package main

import "fmt"

// check if the column is not sorted, we have to remove it
func minDeletionSize(strs []string) int {
	var res int = 0
	var i, j int
	n := len(strs)
	m := len(strs[0])
	sorted := make([]bool, n-1)
	for j = 0; j < m; j++ {
		for i = 0; i < n-1; i++ {
			if !sorted[i] && strs[i][j] > strs[i+1][j] {
				res++
				break
			}
		}
		if i < n-1 {
			continue
		}
		for i = 0; i < n-1; i++ {
			sorted[i] = sorted[i] || strs[i][j] < strs[i+1][j]
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
