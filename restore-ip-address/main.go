package main

import (
	"fmt"
	"strconv"
)

var res []string

func restoreIpAddresses(s string) []string {
	var restored string
	dfs(s, 0, restored, 0)
	return res
}

func dfs(s string, index int, restored string, count int) {
	if count > 4 {
		return
	}
	if count == 4 && index == len(s) {
		res = append(res, restored)
		return
	}

	for i := 1; i < 4; i++ {
		if index+i > len(s) {
			break
		}

		subString := s[index : index+i]
		value, _ := strconv.Atoi(subString)
		if (string(subString[0]) == "0" && len(subString) > 1) || (i == 3 && value > 255) {
			continue
		} else {
			if count == 3 {
				dfs(s, index+i, restored+subString, count+1)
			} else {
				dfs(s, index+i, restored+subString+".", count+1)
			}
		}
	}
}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}
