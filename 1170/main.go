package main

import (
	"fmt"
	"sort"
)

func f(s string) int {
	var count = make([]int, 27)
	for i := 0; i < len(s); i++ {
		count[int(s[i])-int('a')]++
	}
	for i := 0; i < len(count); i++ {
		if count[i] != 0 {
			return count[i]
		}
	}
	return 0
}

func numSmallerByFrequency(queries []string, words []string) []int {
	var ws []int
	for i := 0; i < len(words); i++ {
		ws = append(ws, f(words[i]))
	}

	sort.Slice(ws, func(i, j int) bool {
		return ws[i] < ws[j]
	})
	fmt.Println(ws)

	var res []int
	for i := 0; i < len(queries); i++ {
		fs := f(queries[i])
		idx := sort.Search(len(ws), func(i int) bool {
			return ws[i] > fs
		})

		if idx == len(ws) {
			res = append(res, 0)
			continue
		}
		if ws[idx] == fs {
			res = append(res, len(ws)-idx-1)
		} else {
			res = append(res, len(ws)-idx)
		}
	}
	return res
}

func main() {
	//fmt.Println(numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
	//fmt.Println(numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
	fmt.Println(numSmallerByFrequency([]string{"bba", "abaaaaaa", "aaaaaa", "bbabbabaab", "aba", "aa", "baab", "bbbbbb", "aab", "bbabbaabb"}, []string{"aaabbb", "aab", "babbab", "babbbb", "b", "bbbbbbbbab", "a", "bbbbbbbbbb", "baaabbaab", "aa"}))
}
