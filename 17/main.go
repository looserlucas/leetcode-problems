package main

import (
	"fmt"
	"strings"
)

var letter map[int]string
var result []string

func letterCombinations(digits string) []string {
	result = nil
	letter = make(map[int]string)
	letter[2] = "abc"
	letter[3] = "def"
	letter[4] = "ghi"
	letter[5] = "jkl"
	letter[6] = "mno"
	letter[7] = "pqrs"
	letter[8] = "tuv"
	letter[9] = "wxyz"
	dfs(digits, 0, "")
	return result
}

func dfs(digits string, now int, current string) {
	if now == len(digits) {
		result = append(result, current)
		return
	}
	number := letter[int(digits[now])-int('0')]
	for _, v := range number {
		newString := strings.Clone(current)
		newString = newString + string(v)
		dfs(digits, now+1, newString)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
