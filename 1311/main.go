package main

import (
	"fmt"
	"sort"
)

func greater(a, b string) bool {
	var n int
	if len(a) < len(b) {
		n = len(a)
	} else {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}
	if len(a) > len(b) {
		return true
	} else {
		return false
	}
}

func watchedVideosByFriends(w [][]string, f [][]int, id int, level int) []string {
	var q [][]int
	var done []bool = make([]bool, len(f))
	var wm = make(map[string]int)
	q = append(q, []int{id, 0})
	i := 0
	for i < len(q) {
		if done[q[i][0]] {
			i++
			continue
		} else {
			done[q[i][0]] = true
		}
		if q[i][1] == level {
			for j := 0; j < len(w[q[i][0]]); j++ {
				wm[w[q[i][0]][j]]++
			}
		} else {
			for j := 0; j < len(f[q[i][0]]); j++ {
				if !done[f[q[i][0]][j]] {
					tmp := []int{f[q[i][0]][j], q[i][1] + 1}
					q = append(q, tmp)
				}
			}
		}
		i++
	}
	fmt.Println(wm)

	type wordFreq struct {
		Word string
		Fre  int
	}
	var a []wordFreq
	for k, v := range wm {
		tmp := wordFreq{
			Word: k,
			Fre:  v,
		}
		a = append(a, tmp)
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].Fre < a[j].Fre {
			return true
		} else if a[i].Fre == a[j].Fre {
			return !greater(a[i].Word, a[j].Word)
		}
		return false
	})
	var res []string
	for i := 0; i < len(a); i++ {
		res = append(res, a[i].Word)
	}

	return res
}

func main() {
	/*
		fmt.Println(watchedVideosByFriends([][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}, [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}, 0, 1))
		fmt.Println(watchedVideosByFriends([][]string{{"A", "B"}, {"C"}, {"B", "C"}, {"D"}}, [][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}}, 0, 2))
	*/

	fmt.Println(watchedVideosByFriends([][]string{{"bjwtssmu"}, {"aygr", "mqls"}, {"vrtxa", "zxqzeqy", "nbpl", "qnpl"}, {"r", "otazhu", "rsf"}, {"bvcca", "ayyihidz", "ljc", "fiq", "viu"}}, [][]int{{3, 2, 1, 4}, {0, 4}, {4, 0}, {0, 4}, {2, 3, 1, 0}}, 3, 1))
}
