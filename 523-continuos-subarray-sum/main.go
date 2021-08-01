package main

func checkSubarraySum(nums []int, k int) bool {
	mapper := make(map[int]int)
	sum := 0
	mapper[0] = -1
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if k != 0 {
			sum %= k
		}
		if v, ok := mapper[sum]; ok {
			if i-v > 1 {
				return true
			}
		} else {
			mapper[sum] = i
		}
	}
	return false
}

func main() {
}
