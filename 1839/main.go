package main

import "fmt"

func longestBeautifulSubstring(s string) int {
	count := 0
	res := 0
	var exist = make(map[int]bool)
	for i := 0; i < len(s); i++ {
		if i > 0 && int(s[i]) < int(s[i-1]) {
			_, ok1 := exist[int('a')]
			_, ok2 := exist[int('e')]
			_, ok3 := exist[int('i')]
			_, ok4 := exist[int('o')]
			_, ok5 := exist[int('u')]
			if ok1 && ok2 && ok3 && ok4 && ok5 {
				if count > res {
					res = count
				}
			}
			delete(exist, int('a'))
			delete(exist, int('e'))
			delete(exist, int('i'))
			delete(exist, int('o'))
			delete(exist, int('u'))
			count = 0
		}
		exist[int(s[i])] = true
		count++
	}
	if count > 0 {
		_, ok1 := exist[int('a')]
		_, ok2 := exist[int('e')]
		_, ok3 := exist[int('i')]
		_, ok4 := exist[int('o')]
		_, ok5 := exist[int('u')]
		if ok1 && ok2 && ok3 && ok4 && ok5 {
			if count > res {
				res = count
			}
		}
	}
	return res
}

func main() {
	fmt.Println(longestBeautifulSubstring("aeiaaioaaaaeiiiiouuuooaauuaeiu"))
	fmt.Println(longestBeautifulSubstring("aeeeiiiioooauuuaeiou"))
	fmt.Println(longestBeautifulSubstring("aeeeeeeiiiiiioooooou"))
}
