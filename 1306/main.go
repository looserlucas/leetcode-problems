package main

import "fmt"

var a []int
var memo []int

func canReach(arr []int, start int) bool {
	a = arr
	memo = nil
	done := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		memo = append(memo, -1)
		done[i] = false
	}
	done[start] = true
	return dfs(start, done)
}

func dfs(i int, done map[int]bool) bool {
	if a[i] == 0 {
		return true
	}
	if memo[i] != -1 {
		if memo[i] == 1 {
			return true
		} else {
			return false
		}
	}
	var ok1, ok2 bool = false, false
	if i-a[i] >= 0 && !done[i-a[i]] {
		done[i-a[i]] = true
		ok1 = dfs(i-a[i], done)
		done[i-a[i]] = false
	}
	if i+a[i] < len(a) && !done[i+a[i]] {
		done[i+a[i]] = true
		ok2 = dfs(i+a[i], done)
		done[i+a[i]] = false
	}
	if ok1 || ok2 {
		memo[i] = 1
	} else {
		memo[i] = 0
	}
	return ok1 || ok2
}

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0))
	fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2))
}
