package main

import "fmt"

func frequencySort(s string) string {
	var count = make(map[byte]int)
	var max = -1
	for i := 0; i < len(s); i++ {
		count[s[i]]++
		if count[s[i]] > max {
			max = count[s[i]]
		}
	}
	bucket := make([][]byte, max+1)
	for k, v := range count {
		bucket[v] = append(bucket[v], k)
	}
	var res string
	for i := max; i >= 1; i-- {
		for j := 0; j < len(bucket[i]); j++ {
			for k := 0; k < i; k++ {
				res += string(bucket[i][j])
			}
		}
	}
	return res
}

func main() {
	fmt.Println(frequencySort("tree"))
}
