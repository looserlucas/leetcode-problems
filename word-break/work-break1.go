package main

func wordBreak1(s string, wordDict []string) bool {
	var dp = make(map[int]bool)
	var wordDictMap = make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		wordDictMap[wordDict[i]] = true
	}
	return wordBreakMemo(s, wordDictMap, 0, dp)
}

func wordBreakMemo(s string, wordDictMap map[string]bool, start int, dp map[int]bool) bool {
	if start == len(s) {
		return true
	}

	if val, ok := dp[start]; ok {
		return val
	}

	for end := start + 1; end <= len(s); end++ {
		if _, ok := wordDictMap[s[start:end]]; ok && wordBreakMemo(s, wordDictMap, end, dp) {
			dp[start] = true
			return dp[start]
		}
	}
	dp[start] = false
	return false
}

// bottom up gonna be like
// start will be from len(s) backword
// dp[start] = dp[j] && dp[j:len(s)] (j from start+1 -> len(s))
