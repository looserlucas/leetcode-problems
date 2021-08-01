package main

import "fmt"

func numberOfArithmeticSlices(A []int) int {
	mapper := make(map[int]map[int]int)
	var res int
	for i := 0; i < len(A); i++ {
		mapper[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			if mapper[j][A[i]-A[j]] > 0 {
				res += mapper[j][A[i]-A[j]]
			}
			mapper[i][A[i]-A[j]] += mapper[j][A[i]-A[j]] + 1
		}
	}

	return res
}

func main() {
	A := []int{1, 1, 2, 3}
	fmt.Println(numberOfArithmeticSlices(A))
}
