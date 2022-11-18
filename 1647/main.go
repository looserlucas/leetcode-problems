package main

import "fmt"

func minDeletions(s string) int {
	var freq = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	fmt.Println(freq)
	var bucket = make([]int, len(s)+1)
	for _, v := range freq {
		bucket[v]++
	}
	fmt.Println(bucket)
	var res int = 0
	for i := len(bucket) - 1; i > 0; i-- {
		if bucket[i] > 1 {
			res += bucket[i] - 1
			bucket[i-1] += bucket[i] - 1
		}
	}
	return res
}
func main() {
	fmt.Println(minDeletions("aab"))
}
