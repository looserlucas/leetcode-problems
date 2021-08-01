package main

// k concat max sum problem on Leetcode
const MOD = 1e9 + 7

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func kConcatenationMaxSum(arr []int, k int) int {
	if k == 1 {
		res := 0
		sum1 := 0
		for i := 0; i < len(arr); i++ {
			sum1 += arr[i]
			res = max(res, sum1)
			if sum1 < 0 {
				sum1 = 0
			}
		}

		return res
	}

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	arr = append(arr, arr...)
	res := 0
	sum1 := 0
	for i := 0; i < len(arr); i++ {
		sum1 += arr[i]
		res = max(res, sum1)
		if sum1 < 0 {
			sum1 = 0
		}
	}

	result := max(res, max(sum*k, sum*(k-2)+res))
	return result % MOD
}

func main() {
}
