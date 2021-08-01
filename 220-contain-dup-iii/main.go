package main

import "fmt"

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	w := t + 1
	bucket := make(map[int]int)
	for i, v := range nums {
		if i >= k {
			delete(bucket, nums[i-k]/w)
		}

		var idx int
		if v >= 0 {
			idx = v / w
		} else {
			idx = v/w - 1
		}
		if _, ok := bucket[idx]; ok {
			fmt.Println(i, bucket[idx])
			return true
		} else if _, ok := bucket[idx+1]; ok && abs(v-nums[bucket[idx+1]]) < w {
			fmt.Println(i, bucket[idx+1])
			return true
		} else if _, ok := bucket[idx-1]; ok && abs(v-nums[bucket[idx-1]]) < w {
			fmt.Println(i, bucket[idx-1])
			return true
		}
		bucket[idx] = i
	}
	return false
}

func main() {
	nums := []int{0, 10, 22, 15, 0, 5, 22, 12, 1, 5}
	k := 3
	t := 3
	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t))
}
