package main

// take right most, left=0, right=len(arr)-1 for whole array
func BinarySearchTakeRight(arr []int, left, right int, val int) int {
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
	} else if arr[left] == val {
		return left
	} else {
		return left + 1
	}
}

// take left most, left=0, right=len(arr)-1 for whole array
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
