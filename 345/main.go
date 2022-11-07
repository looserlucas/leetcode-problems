package main

import "fmt"

func reverseVowels(s string) string {
	var vp []int
	var v []byte
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			vp = append(vp, i)
			v = append(v, s[i])
		}
	}
	if len(v) == 0 {
		return s
	}
	var res string
	pos := 0
	vowel := len(v) - 1
	for i := 0; i < len(s); i++ {
		if pos == len(vp) {
			res = res + string(s[i])
		} else {
			if i != vp[pos] {
				res = res + string(s[i])
			} else {
				res = res + string(v[vowel])
				pos++
				vowel--
			}
		}
	}
	return res
}
func main() {
	fmt.Println(reverseVowels("aA"))
}
