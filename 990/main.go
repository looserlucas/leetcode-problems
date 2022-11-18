package main

import "fmt"

var parent []int
var rank []int

func find(i int) int {
	if parent[i] == i {
		return i
	} else {
		res := find(parent[i])
		parent[i] = res
		return res
	}
}

func union(i, j int) {
	irep := find(i)
	jrep := find(j)

	if irep == jrep {
		return
	}
	irank := rank[irep]
	jrank := rank[jrep]

	if irank < jrank {
		parent[irep] = jrep
	} else if jrank < irank {
		parent[jrep] = irep
	} else {
		parent[irep] = jrep
		rank[jrep]++
	}
}

func equationsPossible(equations []string) bool {
	parent = make([]int, 26)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	rank = make([]int, 26)
	for i := 0; i < len(equations); i++ {
		if equations[i][1] == '=' {
			a := int(equations[i][0]) - int('a')
			b := int(equations[i][3]) - int('a')
			union(a, b)
		}
	}
	for i := 0; i < len(equations); i++ {
		if equations[i][1] == '!' {
			a := int(equations[i][0]) - int('a')
			b := int(equations[i][3]) - int('a')
			if find(a) == find(b) {
				return false
			}
		}

	}
	return true
}

func main() {
	fmt.Println(equationsPossible([]string{"a==b", "b!=a"}))
	fmt.Println(equationsPossible([]string{"b==a", "a==b"}))
}
