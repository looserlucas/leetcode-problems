package main

func binarySearch(a []int32, target int32) int32 {
	left := int32(0)
	right := int32(len(a) - 1)
	for left < right {
		mid := left + (right-left+1)/2
		if target > a[mid] {
			right = mid - 1
		} else {
			left = mid
		}
	}
	var res int32
	if a[left] == target {
		res = left + 1
	} else if a[left] > target {
		res = left + 2
	} else {
		if left == 0 {
			res = 1
		} else {
			res = left
		}
	}

	return res
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	var filteredRank []int32
	if len(ranked) == 0 {
		filteredRank = ranked
	} else {
		filteredRank = append(filteredRank, ranked[0])
		for i := range ranked {
			if ranked[i] != filteredRank[len(filteredRank)-1] {
				filteredRank = append(filteredRank, ranked[i])
			}
		}
	}
	var res []int32
	for i := range player {
		res = append(res, binarySearch(filteredRank, player[i]))
	}

	return res
}

func main() {
	climbingLeaderboard([]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120})
}
