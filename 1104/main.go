package main

import "fmt"

func pathInZigZagTree(label int) []int {
	var tmp []int
	level := -1
	n := label
	for n > 0 {
		n = n / 2
		level++
	}
	var end, start = 2, 1
	for i := 2; i <= level; i++ {
		end *= 2
		start *= 2
	}
	fmt.Println(start, end, level, label)
	for label > 0 {
		if end-1 == start {
			tmp = append(tmp, 1)
			label = label / 2
		} else {
			var v int
			tmp = append(tmp, label)
			if level%2 == 0 {
				v = label
			} else {
				fmt.Println(v, level)
				v = start + end - 1 - label
			}
			label = label / 2
			start = start / 2
			end = end / 2
			level--
		}
	}
	var res []int
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}
	return res
}

func main() {
	//fmt.Println(pathInZigZagTree(16))
	fmt.Println(pathInZigZagTree(14))
	//fmt.Println(pathInZigZagTree(26))
}
