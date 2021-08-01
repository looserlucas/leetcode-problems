package main 

// dp[i][j] = ways to generate string t[0:j] from string s[0:i] 
// dp[i][j] = dp[i-1][j-1] + dp[i][j-1] (s[i] = t[j])
//          = dp[i-1][j] (s[i] != t[j])
func numDistinct(s string, t string) int {
	s = " " + s
	t = " " + t
	var dp [][]int
	dp = make([][]int, len(t))
	for i := 0; i < len(t); i++ {
		dp[i] = make([]int, len(s))
		if i == 0 {
			for j := 0; j < len(s); j++ {
				dp[i][j] = 1
			}
		}
	}

	for i := 1; i < len(t); i++ {
		for j := 1; j < len(s); j++ {
			if t[i] == s[j] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(t)-1][len(s)-1]
}

func main() {
}
