package main

import "fmt"

const MODER = uint64(1e9 + 7)
const MAX = uint64(1e11)

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

// take right side
func BinarySearchTakeRight(arr []int, left, right int, val int) int {
	for left < right {
		mid := (right + left) / 2
		if arr[mid] >= val {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	} else {
		return b
	}
}
func minWastedSpace(packages []int, boxes [][]int) int {
	// o(nlogn)
	packages = quickSort(packages, 0, len(packages)-1)
	var pSum uint64
	for i := 0; i < len(packages); i++ {
		pSum += uint64(packages[i])
	}

	var res uint64 = MAX
	// o(m)
	for i := 0; i < len(boxes); i++ {
		// o(mlogm)
		temp := quickSort(boxes[i], 0, len(boxes[i])-1)
		if packages[len(packages)-1] > temp[len(temp)-1] {
			continue
		}
		i := 0
		j := 0
		var sum uint64
		// o(mlogn)
		for b := 0; b < len(temp); b++ {
			j = BinarySearchTakeRight(packages, 0, len(packages), temp[b]+1)
			sum += uint64((j - i)) * uint64(temp[b])
			i = j
		}

		res = min(res, uint64(sum))
	}
	if res != MAX {
		return int((res - pSum) % MODER)
	} else {
		return -1
	}
}

func main() {
	packages := []int{2, 3, 5}
	boxes := [][]int{{4, 8}, {2, 8}}
	fmt.Println(minWastedSpace(packages, boxes))
}
