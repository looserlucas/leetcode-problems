package main

import (
	"fmt"
	"sort"
)

func smaller(a, b string) bool {
	if len(a) < len(b) {
		return true
	} else if len(a) > len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return false
}

func removeSubfolders(folder []string) []string {
	sort.Slice(folder, func(i, j int) bool {
		return smaller(folder[i], folder[j])
	})

	var seen = make(map[string]bool)
	for i := 0; i < len(folder); i++ {
		add := true
		for j := 1; j < len(folder[i]); j++ {
			if folder[i][j] == '/' && seen[folder[i][0:j]] {
				add = false
				break
			}
		}
		if add {
			seen[folder[i]] = true
		}
	}

	var res []string
	for k, _ := range seen {
		res = append(res, k)
	}
	fmt.Println(seen)
	return res
}

func main() {
	fmt.Println(removeSubfolders([]string{"/a", "/c/d", "/c/d/e", "/c/f", "/a/b", "/d", "/b"}))
}
