package main

import (
	"fmt"
)

func rev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/*
func kthPalindrome(queries []int, intLength int) []int64 {
	if intLength == 1 {
		var res []int64
		for i := 0; i < len(queries); i++ {
			if queries[i] > 9 {
				res = append(res, -1)
			} else {
				res = append(res, int64(queries[i]))
			}
		}
		return res
	}
	if intLength == 2 {
		var res []int64
		for i := 0; i < len(queries); i++ {
			if queries[i] > 9 {
				res = append(res, -1)
			} else {
				res = append(res, int64(queries[i]*10+queries[i]))
			}
		}
		return res
	}
	var res []int64
	var cap = make([]int64, intLength+1)
	var base1 = []string{"9", "0", "1", "2", "3", "4", "5", "6", "7", "8"}
	var base2 = []string{"99", "00", "11", "22", "33", "44", "55", "66", "77", "88"}
	cap[1] = 10
	cap[2] = 10
	for i := 3; i <= intLength; i++ {
		cap[i] = cap[i-2] * 9
	}
	for i := 0; i < len(queries); i++ {
		if int64(queries[i]) > cap[intLength] {
			res = append(res, -1)
		} else {
			// this gonna be x...x
			now := intLength
			var x = int64(queries[i])
			var s string
			for now > 0 {
				if now == 1 {
					s = s + base1[x] + rev(s)
					now = 0
				} else if now == 2 {
					s = s + base2[x] + rev(s)
					now = 0
				} else {
					var y = int64(0)
					z := x % cap[now-2]
					if z == 0 {
						y--
					}
					y = y + x/cap[now-2] + 1
					//fmt.Println(x, y)
					s += strconv.Itoa(int(y))
					now -= 2
					x = z
				}
			}
			tmp, _ := strconv.ParseInt(s, 10, 64)
			res = append(res, tmp)
		}
	}
	return res
}
*/
func kthPalindrome(queries []int, intLength int) []int64 {
	
}
func main() {
	fmt.Println(kthPalindrome([]int{1, 2, 3, 4, 5, 10, 11, 30, 31, 20, 21, 90}, 3))
	fmt.Println(kthPalindrome([]int{2, 4, 6, 90}, 4))
}
