package main

import "fmt"

type stepStruct struct {
	value     int
	stepToGet int
}

func minimumJumps(forbidden []int, a int, b int, x int) int {
	var marker [2001]int

	// maxForbidden is the ceiling
	maxForbidden := 0
	for i := 0; i < len(forbidden); i++ {
		marker[forbidden[i]] = -1
		if maxForbidden < forbidden[i] {
			maxForbidden = forbidden[i]
		}
	}
	if x > maxForbidden {
		maxForbidden = x
	}
	maxForbidden = maxForbidden + a + b

	var step []*stepStruct
	stepElement := &stepStruct{
		value:     a,
		stepToGet: 1,
	}

	step = append(step, stepElement)
	for i := 0; i < len(step); i++ {
		if step[i].value > maxForbidden {
			continue
		}
		if step[i].value == x {
			return step[i].stepToGet
		}

		if step[i].value > b && marker[step[i].value-b] == 0 {
			stepElement := &stepStruct{
				value:     step[i].value - b,
				stepToGet: step[i].stepToGet + 1,
			}
			step = append(step, stepElement)
			marker[step[i].value-b] = 1
		}
		if marker[step[i].value+a] == 0 {
			stepElement := &stepStruct{
				value:     step[i].value + a,
				stepToGet: step[i].stepToGet + 1,
			}
			step = append(step, stepElement)
			marker[step[i].value+a] = 1
		}
	}
	return -1
}

func main() {
	forbidden := []int{8, 3, 16, 6, 12, 20}
	a := 15
	b := 13
	x := 11
	fmt.Println(minimumJumps(forbidden, a, b, x))

}
