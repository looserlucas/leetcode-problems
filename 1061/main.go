package main

import (
	"fmt"
	"strings"
)

var parent []int

func find(i int) int {
	if i == parent[i] {
		return i
	} else {
		p := find(parent[i])
		parent[i] = p
		return p
	}
}

func union(i, j int) {
	ip := find(i)
	jp := find(j)

	if ip == jp {
		return
	}

	if ip < jp {
		parent[jp] = ip
	} else {
		parent[ip] = jp
	}
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent = make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	for i := 0; i < len(s1); i++ {
		c1 := int(s1[i]) - int('a')
		c2 := int(s2[i]) - int('a')
		if c1 != c2 {
			union(c1, c2)
		}
	}

	var res strings.Builder
	for i := 0; i < len(baseStr); i++ {
		c := find(int(baseStr[i]) - int('a'))
		res.WriteByte(byte(c + int('a')))
	}
	return res.String()
}

func main() {
	fmt.Println(smallestEquivalentString("parker", "morris", "parser"))
}
