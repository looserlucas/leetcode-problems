package main

import "fmt"

var result []string

func generateParenthesis(n int) []string {
	result = nil
	dfs(n, 0, "")
	return result
}

func clone(a string) string {
	var b string
	for i := 0; i < len(a); i++ {
		b = b + string(a[i])
	}
	return b
}

func dfs(n int, now int, res string) {
	if len(res) == n*2 {
		if now == 0 {
			fmt.Println(res)
			result = append(result, res)
		}
		return
	}
	if now > 0 {
		newRes := clone(res)
		newRes = newRes + ")"
		dfs(n, now-1, newRes)
	}
	if now < n {
		newRes := clone(res)
		newRes = newRes + "("
		dfs(n, now+1, newRes)
	}
}

func main() {
	fmt.Println(generateParenthesis(1))
}
