package main

import "fmt"

func minimumRefill(plant []int, capacityA int, capacityB int) int {
	i := 0
	j := len(plant) - 1
	nowA := capacityA
	nowB := capacityB
	res := 0
	for i <= j {
		fmt.Println(i, j, res)
		if i == j {
			if nowA >= nowB {
				if nowA < plant[i] {
					res++
				}
			} else {
				if nowB < plant[i] {
					res++
				}
			}
			j--
			continue
		}
		if nowA < plant[i] {
			res++
			nowA = capacityA - plant[i]
		} else {
			nowA = nowA - plant[i]
		}
		if nowB < plant[j] {
			res++
			nowB = capacityB - plant[j]
		} else {
			nowB = nowB - plant[j]
		}
		i++
		j--
	}
	return res
}
func main() {
	fmt.Println(minimumRefill([]int{1, 2, 4, 4, 5}, 6, 5))
}
