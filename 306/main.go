package main

import (
	"fmt"
)

func isAdditiveNumber(num string) bool {
	return dfs(num, 0, -1, -1)
}

func dfs(s string, index int, first, second int64) bool {
	if index == len(s) {
		return false
	}
	var val int64 = 0
	if first == -1 {
		for i := index; i < len(s); i++ {
			val = val*10 + int64(int(s[i])-int('0'))
			if i > index && s[index] == '0' || val > int64(1e18) {
				break
			}
			if dfs(s, i+1, val, second) {
				return true
			}
		}
	} else if second == -1 {
		for i := index; i < len(s); i++ {
			val = val*10 + int64(int(s[i])-int('0'))
			if i > index && s[index] == '0' || val > int64(1e18) {
				break
			}
			if dfs(s, i+1, first, val) {
				return true
			}
		}
	} else {
		for i := index; i < len(s); i++ {
			val = val*10 + int64(int(s[i])-int('0'))
			if i > index && s[index] == '0' || val > int64(1e18) {
				break
			}
			if first+second == val {
				if i == len(s)-1 {
					return true
				}
				return dfs(s, i+1, second, val)
			}
			if first+second < val {
				return false
			}
		}
	}
	return false
}

func main() {
	fmt.Println(isAdditiveNumber("112358"))
	fmt.Println(isAdditiveNumber("199100199"))
	fmt.Println(isAdditiveNumber("1023"))
	fmt.Println(isAdditiveNumber("101"))
	fmt.Println(isAdditiveNumber("1202"))
	fmt.Println(isAdditiveNumber("125"))
	fmt.Println(isAdditiveNumber("0000"))
}
