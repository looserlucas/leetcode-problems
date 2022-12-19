package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var s []int
	for _, t := range tokens {
		if t != "*" && t != "/" && t != "+" && t != "-" {
			n, _ := strconv.Atoi(t)
			s = append(s, n)
		} else {
			switch t {
			case "+":
				x := s[len(s)-1]
				y := s[len(s)-2]
				s = s[0 : len(s)-2]
				s = append(s, x+y)
			case "-":
				x := s[len(s)-1]
				y := s[len(s)-2]
				s = s[0 : len(s)-2]
				s = append(s, y-x)
			case "*":
				x := s[len(s)-1]
				y := s[len(s)-2]
				s = s[0 : len(s)-2]
				s = append(s, x*y)
			case "/":
				x := s[len(s)-1]
				y := s[len(s)-2]
				s = s[0 : len(s)-2]
				s = append(s, y/x)
			}
		}
	}
	return s[0]
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
