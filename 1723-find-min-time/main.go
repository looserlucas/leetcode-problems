package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxSlice(a []int) int {
	max := 0
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func minimumTimeRequired(jobs []int, k int) int {
	var worker = make([]int, k)
	var res int = 9999999999
	dfs(jobs, k, 0, worker, &res)
	return res
}

func dfs(jobs []int, k int, cur int, worker []int, res *int) {
	if cur == len(jobs) {
		*res = min(*res, maxSlice(worker))
		return
	}

	var seen = make(map[int]bool)
	for i := 0; i < k; i++ {
		//branch cutting
		if _, ok := seen[worker[i]]; ok {
			continue
		}

		if worker[i]+jobs[cur] >= *res {
			continue
		}

		worker[i] += jobs[cur]
		seen[worker[i]] = true
		dfs(jobs, k, cur+1, worker, res)
		worker[i] -= jobs[cur]
	}

	return
}

func main() {
	jobs := []int{9899456, 8291115, 9477657, 9288480, 5146275, 7697968, 8573153, 3582365, 3758448, 9881935, 2420271, 4542202}
	k := 9
	fmt.Println(minimumTimeRequired(jobs, k))
}
