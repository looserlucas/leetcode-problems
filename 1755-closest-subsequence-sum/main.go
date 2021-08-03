package main

import "fmt"

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// take left side
func BinarySearch(arr []int, left, right int, val int) int {
	for left < right {
		mid := left + (right-left+1)/2
		if arr[mid] > val {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func dfs(i, cur int, a []int, sums *[]int) {
	if i == len(a) {
		*sums = append(*sums, cur)
		return
	}

	dfs(i+1, cur, a, sums)
	dfs(i+1, cur+a[i], a, sums)
}

// main idea is to divide nums into 2 halves: nums1 and nums2, then dfs for sub sequence sub, which will take 2^(n/2) for each, then take sums1[i] and binary search for the sum2[j]-goal
func minAbsDifference(nums []int, goal int) int {
	var sum1, sum2 []int
	nums1 := nums[:len(nums)/2]
	nums2 := nums[len(nums)/2:]

	dfs(0, 0, nums1, &sum1)
	dfs(0, 0, nums2, &sum2)
	ans := int(1e9)
	sum2 = quickSort(sum2, 0, len(sum2)-1)
	fmt.Println(sum2)
	for i := 0; i < len(sum1); i++ {
		j := BinarySearch(sum2, 0, len(sum2), goal-sum1[i])
		if j < len(sum2) {
			ans = min(ans, abs(goal-sum1[i]-sum2[j]))
		}
		/*
			if j > 0 {
				ans = min(ans, abs(goal-sum1[i]-sum2[j-1]))
			}
		*/
	}
	return ans
}

func main() {
	nums := []int{1556913, -259675, -7667451, -4380629, -4643857, -1436369, 7695949, -4357992, -842512, -118463}
	goal := -9681425
	fmt.Println(minAbsDifference(nums, goal))
}
