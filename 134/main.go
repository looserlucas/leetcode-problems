package main

import "fmt"

// the idea is simple:
// if from station i, we can get to station j, then we can stop caring about i->j and start calculating from station j+1 onward.
// this is because i->j is a self substainable route.
// we also care about if sum(gas[i])-sum(cost[i]) < 0 => return -1
func canCompleteCircuit2Pass(gas []int, cost []int) int {
	total := 0
	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
	}
	if total < 0 {
		return -1
	}

	tank := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		if tank+gas[i]-cost[i] < 0 {
			start = i + 1
			tank = 0
		} else {
			tank += gas[i] - cost[i]
		}
	}
	return start
}

func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	tank := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		if tank+gas[i]-cost[i] < 0 {
			start = i + 1
			total += tank + gas[i] - cost[i]
			tank = 0
		} else {
			tank += gas[i] - cost[i]
		}
	}
	if total+tank < 0 {
		return -1
	} else {
		return start
	}
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
