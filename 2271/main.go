package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})

	pre := make([]int, len(tiles))
	pre[0] = tiles[0][1] - tiles[0][0] + 1
	for i := 1; i < len(tiles); i++ {
		pre[i] = pre[i-1] + (tiles[i][1] - tiles[i][0] + 1)
	}

	var res = 0
	for i := 0; i < len(tiles); i++ {
		t := tiles[i][1] - carpetLen + 1
		x := sort.Search(len(tiles), func(i int) bool {
			return tiles[i][1] >= t
		})
		tmp := pre[i] - pre[x] + (tiles[x][1] - max(t, tiles[x][0]) + 1)
		if tmp > res {
			res = tmp
		}
	}
	return res
}

func main() {
	fmt.Println(maximumWhiteTiles([][]int{{1, 5}, {10, 11}, {12, 18}, {20, 25}, {30, 32}}, 10))
	fmt.Println(maximumWhiteTiles([][]int{{1, 7}, {10, 11}, {20, 25}, {30, 32}}, 10))
	fmt.Println(maximumWhiteTiles([][]int{{10, 11}, {1, 1}}, 2))
	fmt.Println(maximumWhiteTiles([][]int{{8051, 8057}, {8074, 8089}, {7994, 7995}, {7969, 7987}, {8013, 8020}, {8123, 8139}, {7930, 7950}, {8096, 8104}, {7917, 7925}, {8027, 8035}, {8003, 8011}}, 9854))
}
