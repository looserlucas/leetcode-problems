package main

import "fmt"

var l, r []int

func countHighestScoreNodes(p []int) int {
	l = make([]int, len(p))
	r = make([]int, len(p))
	tree := make(map[int][]int)
	for i := 1; i < len(p); i++ {
		if _, ok := tree[p[i]]; !ok {
			var temp = []int{i}
			tree[p[i]] = temp
		} else {
			tree[p[i]] = append(tree[p[i]], i)
		}
	}
	dfs(tree, 0)
	fmt.Println(tree)
	fmt.Println(l)
	fmt.Println(r)
	res := -1
	count := 0
	for i := 0; i < len(p); i++ {
		// this is a root
		if i == 0 {
			var temp int
			if l[i] == 0 || r[i] == 0 {
				temp = len(p) - 1
			} else {
				temp = l[i] * r[i]
			}
			if temp > res {
				res = temp
				count = 1
			}
			continue
		}
		// this is a leaf
		if l[i] == 0 && r[i] == 0 {
			temp := len(p) - 1
			if temp > res {
				res = temp
				count = 1
			} else if temp == res {
				count++
			}
			continue
		}
		// this is in the middle
		if l[i] == 0 && r[i] != 0 {
			temp := r[i] * (len(p) - 1 - r[i])
			if temp > res {
				res = temp
				count = 1
			} else if temp == res {
				count++
			}
			continue
		}
		if l[i] != 0 && r[i] == 0 {
			temp := l[i] * (len(p) - 1 - l[i])
			if temp > res {
				res = temp
				count = 1
			} else if temp == res {
				count++
			}
			continue
		}
		temp := l[i] * r[i] * (len(p) - 1 - l[i] - r[i])
		if temp > res {
			res = temp
			count = 1
		} else if temp == res {
			count++
		}
		continue
	}
	return count
}

func dfs(tree map[int][]int, now int) int {
	fmt.Println("in here")
	v, ok := tree[now]
	if !ok {
		// this is a leave
		return 1
	}

	if len(v) == 2 {
		// 2 childrens
		left := dfs(tree, v[0])
		right := dfs(tree, v[1])
		l[now] = left
		r[now] = right
		return left + right + 1
	} else {
		// 1 childrens
		left := dfs(tree, v[0])
		l[now] = left
		return left + 1
	}
}

func main() {
	fmt.Println(countHighestScoreNodes([]int{-1, 2, 0, 2, 0}))
	fmt.Println(countHighestScoreNodes([]int{-1, 2, 0}))
}
