package main

import "fmt"

func wordBreak2(s string, wordDict []string) []string {
	var dp = make(map[string][]string)
	return dfs(s, wordDict, dp)
}

func dfs(s string, wordDict []string, dp map[string][]string) []string {
	// dp[s] is the memo to remember ways that string s can be broken into
	if _, ok := dp[s]; ok {
		return dp[s]
	}

	var res []string
	if len(s) == 0 {
		res = append(res, "")
		return res
	}

	for i := 0; i < len(wordDict); i++ {
		if len(wordDict[i]) > len(s) {
			continue
		}
		if s[0:len(wordDict[i])] == wordDict[i] {
			subList := dfs(s[len(wordDict[i]):], wordDict, dp)
			for j := 0; j < len(subList); j++ {
				if len(subList[j]) == 0 {
					res = append(res, wordDict[i]+subList[j])
				} else {
					res = append(res, wordDict[i]+" "+subList[j])
				}
			}
		}
	}
	dp[s] = res
	return res
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak2(s, wordDict))
}
