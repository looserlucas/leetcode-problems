package main

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

var memo = make(map[int]int)

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}

	if _, ok := memo[n]; !ok {
		if n%2 == 0 {
			memo[n] = 1 + integerReplacement(n/2)
		} else {
			memo[n] = 1 + min(integerReplacement((n+1)), integerReplacement((n-1)))
		}
	}

	return memo[n]
}

func main() {
}
