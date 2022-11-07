package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	pivot := rand.Intn(len(nums))
	return pivot
}

func main() {
	fmt.Println(findKthLargest([]int{1, 2}, 1))
}
