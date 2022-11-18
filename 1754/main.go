package main

import "fmt"

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

func largestMerge(word1 string, word2 string) string {
	i := 0
	j := 0
	var res string
	for i < len(word1) || j < len(word2) {
		fmt.Println(word1[i:], word2[j:], res)
		if i == len(word1) {
			res += string(word2[j])
			j++
			continue
		}
		if j == len(word2) {
			res += string(word1[i])
			i++
			continue
		}
		if word1[i] > word2[j] {
			res += string(word1[i])
			i++
		} else if word1[i] == word2[j] {
			if greater(word1[i:], word2[j:]) {
				res += string(word1[i])
				i++
			} else {
				res += string(word2[j])
				j++
			}
		} else {
			res += string(word2[j])
			j++
		}
	}
	return res
}

func main() {
	/*
		fmt.Println(largestMerge("cabaa", "bcaaa"))
		fmt.Println(largestMerge("abcabc", "abdcaba"))
		fmt.Println(largestMerge("nnnnpnnennenpnnnnneenpnn", "nnnennnnnnpnnennnnennnnee"))
	*/
	fmt.Println(largestMerge("guguuuuuuuuuuuuuuguguuuuguug", "gguggggggguuggguugggggg"))
	fmt.Println(greater("g", "gguggggggguuggguugggggg"))
}
