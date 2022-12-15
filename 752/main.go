package main

import (
	"fmt"
	"strconv"
)

func openLock(deadends []string, target string) int {
	deadendsMap := make(map[string]bool)
	for i := 0; i < len(deadends); i++ {
		deadendsMap[deadends[i]] = true
	}
	if deadendsMap["0000"] {
		return -1
	}
	type per struct {
		state string
		step  int
	}
	var q []per
	q = append(q, per{state: "0000", step: 0})
	deadendsMap["0000"] = true
	for len(q) > 0 {
		if q[0].state == target {
			return q[0].step
		}
		s := q[0].state
		step := q[0].step
		for i := 0; i < 4; i++ {
			x := int(s[i]) - int('0')
			for diff := -1; diff <= 1; diff += 2 {
				y := (x + diff + 10) % 10
				new := s[0:i] + strconv.Itoa(y) + s[i+1:]
				if _, ok := deadendsMap[new]; !ok {
					deadendsMap[new] = true
					q = append(q, per{state: new, step: step + 1})
				}
			}
		}
		q = q[1:]
	}
	return -1
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}
