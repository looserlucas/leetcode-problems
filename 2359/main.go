package main

import (
	"fmt"
)

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	if node1 == node2 {
		return node1
	}
	var done = make([]int, len(edges))
	type node struct {
		Value  int
		From   int
		Degree int
	}

	var q []node
	tmp1 := node{
		Value:  node1,
		From:   1,
		Degree: 0,
	}
	tmp2 := node{
		Value:  node2,
		From:   2,
		Degree: 0,
	}
	done[node1] = 1
	done[node2] = 2
	var i = 0
	found := false
	lastFound := -1
	curDeg := -1
	q = append(q, tmp1)
	q = append(q, tmp2)
	for i < len(q) {
		next := edges[q[i].Value]
		if next == -1 {
			i++
			continue
		}
		if done[next] == 0 {
			tmp := node{
				Value:  next,
				From:   q[i].From,
				Degree: q[i].Degree + 1,
			}
			done[next] = q[i].From
			q = append(q, tmp)
			i++
			continue
		}
		if done[next] != q[i].From {
			if !found {
				curDeg = q[i].Degree + 1
				lastFound = next
				found = true
				i++
			} else {
				nowDeg := q[i].Degree + 1
				if nowDeg > curDeg {
					return lastFound
				} else {
					if next < lastFound {
						return next
					} else {
						return lastFound
					}
				}
			}
		} else {
			i++
			continue
		}
	}
	if found {
		return lastFound
	} else {
		return -1
	}
}

func main() {
	/*
		fmt.Println(closestMeetingNode([]int{2, 2, 3, -1}, 0, 1))
		fmt.Println(closestMeetingNode([]int{1, 2, -1}, 0, 2))
	*/
	fmt.Println(closestMeetingNode([]int{4, 4, 8, -1, 9, 8, 4, 4, 1, 1}, 5, 6))
}
