package main

// take right side
func BinarySearchTakeRight(arr []int, left, right int, val int) int {
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] >= val {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
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
	return left
}
