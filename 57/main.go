package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func insert(a [][]int, new []int) [][]int {
	var res [][]int
	// handle left
	var i = 0
	for i < len(a) {
		if a[i][1] < new[0] {
			res = append(res, a[i])
			i++
		} else {
			break
		}
	}
	// try to merge
	for i < len(a) {
		if a[i][0] <= new[1] && new[1] <= a[i][1] {
			new[0] = min(new[0], a[i][0])
			new[1] = max(new[1], a[i][1])
			i++
			continue
		}
		if a[i][0] <= new[0] && new[0] <= a[i][1] {
			new[0] = min(new[0], a[i][0])
			new[1] = max(new[1], a[i][1])
			i++
			continue
		}
		if a[i][0] >= new[0] && a[i][1] <= new[1] {
			i++
			continue
		}
		break
	}
	res = append(res, new)
	for i < len(a) {
		res = append(res, a[i])
		i++
	}
	return res
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}
