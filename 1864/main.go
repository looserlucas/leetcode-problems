package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minSwaps(s string) int {
	count1 := 0
	count0 := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count1++
		} else {
			count0++
		}
	}
	if count0 > count1+1 || count1 > count0+1 {
		return -1
	}
	res := 0
	now0 := 0
	now1 := 0
	if count1 == count0+1 {
		//101 case
		for i := 0; i < len(s); i++ {
			if i%2 == 0 {
				if s[i] == '0' {
					if now1 == 0 {
						res++
						now0++
					} else {
						now1--
					}
				}
			} else {
				if s[i] == '1' {
					if now0 == 0 {
						res++
						now1++
					} else {
						now0--
					}
				}
			}
		}
	} else if count0 == count1+1 {
		//010 case
		for i := 0; i < len(s); i++ {
			if i%2 == 0 {
				if s[i] == '1' {
					if now0 == 0 {
						res++
						now1++
					} else {
						now0--
					}
				}
			} else {
				if s[i] == '0' {
					if now1 == 0 {
						res++
						now0++
					} else {
						now1--
					}
				}
			}
		}
	} else if count1 == count0 {
		var res1, res2 int
		//101 case
		for i := 0; i < len(s); i++ {
			if i%2 == 0 {
				if s[i] == '0' {
					if now1 == 0 {
						res1++
						now0++
					} else {
						now1--
					}
				}
			} else {
				if s[i] == '1' {
					if now0 == 0 {
						res1++
						now1++
					} else {
						now0--
					}
				}
			}
		}
		//010 case
		for i := 0; i < len(s); i++ {
			if i%2 == 0 {
				if s[i] == '1' {
					if now0 == 0 {
						res2++
						now1++
					} else {
						now0--
					}
				}
			} else {
				if s[i] == '0' {
					if now1 == 0 {
						res2++
						now0++
					} else {
						now1--
					}
				}
			}
		}
		res = min(res1, res2)
	}
	return res
}

func main() {
	fmt.Println(minSwaps("111000"))
	fmt.Println(minSwaps("1111100000"))
	fmt.Println(minSwaps("010"))
	fmt.Println(minSwaps("10"))
	fmt.Println(minSwaps("01"))
	fmt.Println(minSwaps("1110"))
}
