package main

import "fmt"

//https://leetcode.com/problems/scramble-string/
// dp[i][j][k] = if it's possilbe to 'turn i->i+k-1 to j->j+k-1
// 			substring 1   			       substring 2
// => dp[i][j][k] = (dp[i][j][l] && dp[i+l][j+l][k-l]) || (dp[i][j+k-l][l] && dp[i+l][j][k-l])
//                     x1 == y1         x2 == y2             x1 == y2           x2 == y1
// x1 + x2 = x = i->i+k-1
// y1 + y2 = y = j->j+k-1
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var dp [30][30][31]bool
	for k := 1; k <= len(s1); k++ {
		for i := 0; i+k <= len(s1); i++ {
			for j := 0; j+k <= len(s1); j++ {
				if k == 1 {
					dp[i][j][k] = s1[i] == s2[j]
				} else {
					for l := 1; l < k && !dp[i][j][k]; l++ {
						dp[i][j][k] = (dp[i][j][l] && dp[i+l][j+l][k-l]) || (dp[i][j+k-l][l] && dp[i+l][j][k-l])
					}
				}
			}
		}
	}

	return dp[0][0][len(s1)]
}

func main() {
	s1 := "a"
	s2 := "a"
	fmt.Println(isScramble(s1, s2))
}
