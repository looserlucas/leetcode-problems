package main

import "fmt"

func minSwapsCouples(row []int) int {
	var pos = make([]int, len(row))
	for i := 0; i < len(row); i++ {
		pos[row[i]] = i
	}

	res := 0
	// get all i = 0,2,4,... With each i, check if row[i] is even or odd.
	// if row[i] == even => row[i+1] must be row[i]+1 -> swap row[i+1] with row[i]+1
	// if row[i] == odd => row[i+1] must be row[i]-1 -> swap row[i+1] with row[i]-1
	for i := 0; i < len(row)-1; i += 2 {
		if row[i]%2 == 0 && row[i+1] != row[i]+1 {
			idx1 := pos[row[i]+1]
			pos[row[i]+1] = i + 1
			pos[row[i+1]] = idx1
			row[idx1], row[i+1] = row[i+1], row[idx1]
			res++
		} else if row[i]%2 == 1 && row[i+1] != row[i]-1 {
			idx1 := pos[row[i]-1]
			pos[row[i]-1] = i + 1
			pos[row[i+1]] = idx1
			row[idx1], row[i+1] = row[i+1], row[idx1]
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(minSwapsCouples([]int{0, 2, 1, 3}))
	fmt.Println(minSwapsCouples([]int{3, 2, 0, 1}))
}
