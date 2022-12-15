package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxConsecutiveAnswers(a string, k int) int {
	return max(helper(a, 'T', k), helper(a, 'F', k))
}

func helper(a string, b byte, k int) int {
	j := 0
	res := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b {
			k--
		}

		for j < i && k < 0 {
			if a[j] != b {
				k++
			}
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}

func main() {
	fmt.Println(maxConsecutiveAnswers("TTFF", 2))
	fmt.Println(maxConsecutiveAnswers("TFFT", 1))
	fmt.Println(maxConsecutiveAnswers("TTFTTFTT", 1))
}
