package main

import (
	"fmt"
	"sort"
)

func less(a, b string) bool {
	for i := 0; i < len(a); i++ {
		if int(a[i]) < int(b[i]) {
			return true
		} else if int(a[i]) > int(b[i]) {
			return false
		}
	}
	return false
}

var res []string

func findItinerary(tickets [][]string) []string {
	var adj = make(map[string][]string)
	res = nil
	for i := 0; i < len(tickets); i++ {
		_, ok := adj[tickets[i][0]]
		if !ok {
			var tmp []string
			tmp = append(tmp, tickets[i][1])
			adj[tickets[i][0]] = tmp
		} else {
			adj[tickets[i][0]] = append(adj[tickets[i][0]], tickets[i][1])
		}
	}
	for _, v := range adj {
		sort.Slice(v, func(i, j int) bool {
			return !less(v[i], v[j])
		})
	}
	dfs(adj, "JFK")

	var finalRes []string
	for i := len(res) - 1; i >= 0; i-- {
		finalRes = append(finalRes, res[i])
	}
	return finalRes
}

func dfs(adj map[string][]string, now string) {
	for len(adj[now]) > 0 {
		x := adj[now][len(adj[now])-1]
		adj[now] = adj[now][:len(adj[now])-1]
		dfs(adj, x)
	}
	res = append(res, now)
}

func main() {
	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
	fmt.Println(findItinerary([][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}))
}
