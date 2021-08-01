package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	var ok [200][200]bool
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				ok[i][j] = true
			} else if i == 0 {
				ok[i][j] = ok[i][j-1] && (s2[j-1] == s3[j-1])
			} else if j == 0 {
				ok[i][j] = ok[i-1][j] && (s1[i-1] == s3[i-1])
			} else {
				ok[i][j] = (ok[i-1][j] && s1[i-1] == s3[i+j-1]) || (ok[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return ok[len(s1)][len(s2)]
}

func main() {
}
