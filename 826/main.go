package main

import (
	"fmt"
	"sort"
)

func maxProfitAssignment(d []int, p []int, worker []int) int {
	var difPro [][]int
	for i := 0; i < len(d); i++ {
		temp := make([]int, 2)
		temp[0] = d[i]
		temp[1] = p[i]
		difPro = append(difPro, temp)
	}

	sort.SliceStable(difPro, func(i, j int) bool {
		if difPro[i][0] < difPro[j][0] {
			return true
		} else if difPro[i][0] == difPro[j][0] {
			return difPro[i][1] < difPro[j][1]
		}
		return false
	})

	var maxHere = make([]int, len(d))
	cur := -1
	for i := 0; i < len(difPro); i++ {
		if difPro[i][1] > cur {
			cur = difPro[i][1]
		}
		maxHere[i] = cur
	}
	var res int = 0
	for i := 0; i < len(worker); i++ {
		x := worker[i]
		j := sort.Search(len(difPro), func(i int) bool {
			return difPro[i][0] > x
		})
		if j == len(difPro) {
			res += maxHere[len(difPro)-1]
			//fmt.Println(j, maxHere[len(difPro)-1])
			continue
		}
		if difPro[j][0] > x {
			if j > 0 {
				res += maxHere[j-1]
				//fmt.Println(j, maxHere[j-1])
			}
			continue
		}
		//fmt.Println(j, maxHere[j])
		res += maxHere[j]
	}
	fmt.Println(res)
	return res
}

func main() {
	maxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7})
	maxProfitAssignment([]int{85, 47, 57}, []int{24, 66, 99}, []int{40, 25, 25})
	maxProfitAssignment([]int{68, 35, 52, 47, 86}, []int{67, 17, 1, 81, 3}, []int{92, 10, 85, 84, 82})

	maxProfitAssignment([]int{23, 30, 35, 35, 43, 46, 47, 81, 83, 98}, []int{8, 11, 11, 20, 33, 37, 60, 72, 87, 95}, []int{95, 46, 47, 97, 11, 35, 99, 56, 41, 92})
}
