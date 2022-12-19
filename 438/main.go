package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var countP = make(map[byte]int)
	for i := 0; i < len(p); i++ {
		countP[p[i]]++
	}

	var count = make(map[byte]int)
	j := 0
	var res []int
	for i := 0; i < len(s); i++ {
		count[s[i]]++
		for i-j+1 > len(p) {
			count[s[j]]--
			if count[s[j]] == 0 {
				delete(count, s[j])
			}
			j++
		}
		if i-j+1 == len(p) {
			var ok = true
			for k, v := range count {
				if v != countP[k] {
					ok = false
					break
				}
			}
			if ok {
				res = append(res, j)
			}
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
