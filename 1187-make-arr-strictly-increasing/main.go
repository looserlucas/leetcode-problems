package main

import "fmt"

const MAX_INT = int(1e9)
const MIN_INT = -int(1e9)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func BinarySearchTakeRight(arr []int, left, right int, val int) int {
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] >= val {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if arr[left] == val {
		left++
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

// dp[i] is the minimum change to make 0 -> i a strictly increasing array
// like longest increase subarray, dp[i] = min(dp[j]+change(i,j)) (j is from 0 to i-1)
func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	arr2 = quickSort(arr2, 0, len(arr2)-1)
	var tmp []int
	// get rid of duplicate value in arr2
	for i := 0; i < len(arr2); i++ {
		if i+1 < len(arr2) && arr2[i] == arr2[i+1] {
			continue
		}
		tmp = append(tmp, arr2[i])
	}
	arr2 = tmp

	// turn a1 to MIN_INT .... MAX_INT
	tmp = nil
	tmp = append(tmp, MIN_INT)
	tmp = append(tmp, arr1...)
	tmp = append(tmp, MAX_INT)
	arr1 = tmp

	var dp []int
	for i := 0; i < len(arr1); i++ {
		dp = append(dp, MAX_INT)
	}

	dp[0] = 0
	for i := 1; i < len(arr1); i++ {
		for j := 0; j < i; j++ {
			if arr1[j] < arr1[i] && dp[j] != MAX_INT {
				change := check(arr1, arr2, j, i)

				if change >= 0 {
					// because of this min function, being greedy is ok
					dp[i] = min(dp[i], dp[j]+change)
				}
			}
		}
	}
	if dp[len(arr1)-1] == MAX_INT {
		return -1
	} else {
		return dp[len(arr1)-1]
	}
}

// being greedy, check if we can change the whole start+1 -> end-1 with number from arr2
func check(arr1, arr2 []int, start, end int) int {
	if start+1 == end {
		return 0
	}
	min_val := arr1[start]
	max_val := arr1[end]

	idx := BinarySearchTakeRight(arr2, 0, len(arr2)-1, min_val)
	if idx == len(arr2) || arr2[idx] < min_val {
		return -1
	}

	var maxcount = end - start - 1
	endI := idx + maxcount - 1
	if endI < len(arr2) && arr2[endI] < max_val {
		return maxcount
	} else {
		return -1
	}
}

func main() {
	arr1 := []int{1, 5, 3, 6, 7}
	arr2 := []int{1, 6, 3, 3}
	fmt.Println(makeArrayIncreasing(arr1, arr2))
}
