package main

import "fmt"

// * is obstacle
// . is nothing
// # is stone
func rotateTheBox(box [][]byte) [][]byte {
	n := len(box)
	for i := 0; i < n; i++ {
		box[i] = append(box[i], '*')
	}

	m := len(box[0])
	var res = make([][]byte, m)
	for i := 0; i < m; i++ {
		tmp := make([]byte, n)
		res[i] = tmp
	}

	for i := 0; i < n; i++ {
		start := 0
		countSpace := 0
		var rowRes []byte
		for j := 0; j < m; j++ {
			if box[i][j] == '*' {
				var tmp []byte
				for k := 0; k < countSpace; k++ {
					tmp = append(tmp, '.')
				}
				for k := countSpace; k < j-start; k++ {
					tmp = append(tmp, '#')
				}
				tmp = append(tmp, '*')
				rowRes = append(rowRes, tmp...)
				start = j + 1
				countSpace = 0
			} else if box[i][j] == '.' {
				countSpace++
			}
		}
		printString(rowRes)
		for j := 0; j < m; j++ {
			res[j][n-1-i] = rowRes[j]
		}
	}
	return res[1:]
}
func printString(s []byte) {
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), " ")
	}
	fmt.Println()
}

func main() {
	//fmt.Println(rotateTheBox([][]byte{{'#', '.', '#'}}))
	fmt.Println(rotateTheBox([][]byte{{'#', '.', '*', '.'},
		{'#', '#', '*', '.'}}))
}
