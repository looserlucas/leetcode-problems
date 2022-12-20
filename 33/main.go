package main

import "fmt"

// the idea is to first find the rotated spot. We do that by divide and conquer.
// divide the array into 2 parts, by a mid index [lo...mid...hi].
// If a[mid] < a[hi] => the rotated spot is in lo -> mid
// If a[mid] > a[hi] => the rotated spot is in mid+1 -> hi
func search(nums []int, target int) int {
	n := len(nums)
	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] < nums[hi] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	rot := lo

	hi = n - 1
	lo = rot
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > target {
			hi = mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] == target {
			lo = mid
			break
		}
	}
	if nums[lo] == target {
		return lo
	}

	if rot > 0 {
		hi = rot - 1
		lo = 0
		for lo < hi {
			mid := (lo + hi) / 2
			if nums[mid] > target {
				hi = mid
			} else if nums[mid] < target {
				lo = mid + 1
			} else if nums[mid] == target {
				lo = mid
				break
			}
		}
		if nums[lo] == target {
			return lo
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 1))
}
