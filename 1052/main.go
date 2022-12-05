package main

import "fmt"

func maxSatisfied(customers []int, grumpy []int, k int) int {
	normal := make([]int, len(customers))
	notMad := make([]int, len(customers))
	if grumpy[0] == 0 {
		normal[0] = customers[0]
	}
	notMad[0] = customers[0]
	for i := 1; i < len(customers); i++ {
		if grumpy[i] == 0 {
			normal[i] = normal[i-1] + customers[i]
		} else {
			normal[i] = normal[i-1]
		}
		notMad[i] = notMad[i-1] + customers[i]
	}
	fmt.Println(notMad)
	fmt.Println(normal)
	var res int = 0
	n := len(customers) - 1
	for i := 0; i < len(customers); i++ {
		if i-k < 0 {
			tmp := notMad[i] + (normal[n] - normal[i])
			if res < tmp {
				res = tmp
			}
		} else {
			tmp := normal[i-k] + (notMad[i] - notMad[i-k]) + (normal[n] - normal[i])
			if res < tmp {
				res = tmp
			}
		}
	}
	return res
}

func main() {
	fmt.Println(maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
	fmt.Println(maxSatisfied([]int{1}, []int{1}, 1))
	fmt.Println(maxSatisfied([]int{4, 10, 10}, []int{1, 1, 0}, 2))

	fmt.Println(maxSatisfied([]int{3, 8, 8, 7, 1}, []int{1, 1, 1, 1, 1}, 3))
	fmt.Println(maxSatisfied([]int{4, 10, 10}, []int{1, 1, 0}, 2))
}
