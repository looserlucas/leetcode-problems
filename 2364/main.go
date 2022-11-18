package main

import "fmt"

var MOD = 1000000007

func calRes(a int) int {
	return int(int64(a) * int64((a - 1)) / 2)
}

func countBadPairs(nums []int) int64 {
	map1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		a := i - nums[i]
		map1[a] = (map1[a] + 1) % MOD
	}

	var res int
	for _, v := range map1 {
		if v != 1 {
			res = (res + calRes(v)) % MOD
		}
	}

}
func main() {
	fmt.Println("vim-go")
}
