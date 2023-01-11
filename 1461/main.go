package main

import "fmt"

func hasAllCodes(s string, k int) bool {
	target := 1
	lastBit := 1
	for i := 0; i < k; i++ {
		if i != 0 {
			lastBit *= 2
		}
		target *= 2
	}
	if len(s) < target {
		return false
	}
	var count = make(map[int]bool)
	now := 0
	for i := 0; i < k; i++ {
		if s[i] == '0' {
			now = (now << 1)
		} else {
			now = (now << 1)
			now |= 1
		}
	}
	count[now] = true
	for i := k; i < len(s); i++ {
		if now >= lastBit {
			now = now - lastBit
		}
		if s[i] == '0' {
			now = now << 1
		} else {
			now = now << 1
			now |= 1
		}
		count[now] = true
	}
	return len(count) == target
}

func main() {
	fmt.Println(hasAllCodes("00110110", 2))
	fmt.Println(hasAllCodes("0110", 1))
	fmt.Println(hasAllCodes("0110", 2))
}
