package main

import (
	"strconv"
	"strings"
)

func Calculate(s string, iTw map[string]string) string {
	if len(s) != 3 {
		if len(s) == 2 {
			if string(s[0]) == "1" && string(s[1]) == "0" {
				return iTw["10"]
			}
			if string(s[0]) == "1" && string(s[1]) == "1" {
				return iTw["11"]
			}
			if string(s[0]) == "1" && string(s[1]) == "2" {
				return iTw["12"]
			}
			if string(s[0]) == "1" && string(s[1]) == "4" {
				return iTw["4"] + iTw["1*"]
			}

			if string(s[0]) == "1" {
				return iTw[string(s[1])+"th"] + iTw["1*"]
			}

			if string(s[1]) == "0" {
				return iTw[string(s[0])+"th"] + iTw["*0"]
			} else {
				return iTw[string(s[0])+"th"] + iTw["*0"] + " " + iTw[string(s[1])]
			}
		} else {
			return iTw[s]
		}
	} else {
		var res string
		if string(s[0]) != "0" {
			res = res + iTw[string(s[0])] + " " + iTw["*00"]
		}
		if string(s[1]) != "0" {
			if string(s[1]) == "1" && string(s[2]) == "0" {
				return res + " " + iTw["10"]
			}
			if string(s[1]) == "1" && string(s[2]) == "1" {
				return res + " " + iTw["11"]
			}
			if string(s[1]) == "1" && string(s[2]) == "2" {
				return res + " " + iTw["12"]
			}
			if string(s[1]) == "1" && string(s[2]) == "4" {
				return res + " " + iTw["4"] + iTw["1*"]
			}

			if string(s[1]) == "1" {
				return res + " " + iTw[string(s[2])+"th"] + iTw["1*"]
			}

			if string(s[2]) == "0" {
				return res + " " + iTw[string(s[1])+"th"] + iTw["*0"]
			} else {
				return res + " " + iTw[string(s[1])+"th"] + iTw["*0"] + " " + iTw[string(s[2])]
			}
		} else {
			if string(s[2]) == "0" {
				return res
			} else {
				return res + " " + iTw[string(s[2])]
			}
		}
	}
}
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	var iTw map[string]string = make(map[string]string)

	// populate map
	iTw["0"] = ""
	iTw["1"] = "One"
	iTw["10"] = "Ten"
	iTw["11"] = "Eleven"

	iTw["2"] = "Two"
	iTw["12"] = "Twelve"
	iTw["2th"] = "Twen"

	iTw["3"] = "Three"
	iTw["3th"] = "Thir"

	// Fourteen = itw[4] + itw[1*]
	iTw["4"] = "Four"
	iTw["4th"] = "For"

	iTw["5"] = "Five"
	iTw["5th"] = "Fif"

	iTw["6"] = "Six"
	iTw["6th"] = "Six"

	iTw["7"] = "Seven"
	iTw["7th"] = "Seven"

	iTw["8"] = "Eight"
	iTw["8th"] = "Eigh"

	iTw["9"] = "Nine"
	iTw["9th"] = "Nine"

	iTw["*0"] = "ty"
	iTw["1*"] = "teen"

	iTw["*00"] = "Hundred"

	s := strconv.Itoa(num)
	r := len(s) % 3
	d := len(s) / 3

	var res []string
	if r != 0 {
		res = append(res, Calculate(s[:r], iTw))
	}

	for i := 0; i < d; i++ {
		res = append(res, Calculate(s[r:r+3], iTw))
		r = r + 3
	}

	var indicator map[int]string = make(map[int]string)
	indicator[1] = "Thousand"
	indicator[2] = "Million"
	indicator[3] = "Billion"

	var finalRes string
	for i := 0; i < len(res); i++ {
		if res[i] != "" {
			finalRes = finalRes + res[i] + " " + indicator[len(res)-i-1] + " "
		} else {
			finalRes = finalRes + res[i] + " "
		}
	}
	words := strings.Fields(finalRes)
	finalRes = ""
	for i := 0; i < len(words); i++ {
		if i != len(words)-1 {
			finalRes = finalRes + words[i] + " "
		} else {
			finalRes = finalRes + words[i]
		}
	}

	return finalRes
}

func main() {
	numberToWords(1000001)
}
