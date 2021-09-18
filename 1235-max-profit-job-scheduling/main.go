package main

import "fmt"

func partition(arr1, arr, arr2 []int, low, high int) ([]int, []int, []int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			arr1[i], arr1[j] = arr1[j], arr1[i]
			arr2[i], arr2[j] = arr2[j], arr2[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	arr1[i], arr1[high] = arr1[high], arr1[i]
	arr2[i], arr2[high] = arr2[high], arr2[i]
	return arr1, arr, arr2, i
}

func quickSort(arr1, arr, arr2 []int, low, high int) ([]int, []int, []int) {
	if low < high {
		var p int
		arr1, arr, arr2, p = partition(arr1, arr, arr2, low, high)
		arr1, arr, arr2 = quickSort(arr1, arr, arr2, low, p-1)
		arr1, arr, arr2 = quickSort(arr1, arr, arr2, p+1, high)
	}
	return arr1, arr, arr2
}

func BinarySearchTakeLeft(arr []int, left, right int, val int) int {
	for left < right {
		mid := left + (right-left+1)/2
		if arr[mid] > val {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if arr[left] > val {
		return -1
	} else {
		return left
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Basic 0-1 Knapsack problem
// dp[i] = maximum profit till i, so either dp[i] = dp[i-1] (not picking job[j])
//                                      or  dp[i] = dp[j] + profit[i] (j is the maximum profit which jobs[j].end < jobs[i].start, can find this by binary search)
func jobScheduling(start []int, end []int, profit []int) int {
	start, end, profit = quickSort(start, end, profit, 0, len(end)-1)
	var dp = make([]int, len(end))
	dp[0] = profit[0]
	for i := 1; i < len(end); i++ {
		j := BinarySearchTakeLeft(end, 0, len(end)-1, start[i])
		if j != -1 {
			dp[i] = max(dp[i-1], dp[j]+profit[i])
		} else {
			dp[i] = max(dp[i-1], profit[i])
		}
	}
	return dp[len(end)-1]
}

func main() {
	/*
		start := []int{1, 2, 3, 3}
		end := []int{3, 4, 5, 6}
		profit := []int{50, 10, 40, 70}
	*/
	start := []int{1, 2, 3, 4, 6}
	end := []int{3, 5, 10, 6, 9}
	profit := []int{20, 20, 100, 70, 60}
	/*
		start := []int{1}
		end := []int{2}
		profit := []int{5}
	*/

	fmt.Println(jobScheduling(start, end, profit))
}
