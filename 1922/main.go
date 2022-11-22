package main

import "fmt"

var MODN = uint64(1e9 + 7)

func powerMOD(a, b uint64) uint64 {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a % MODN
	}
	p := powerMOD(a, b/2)
	if b%2 == 0 {
		return (p * p) % MODN
	} else {
		return (p * p) % MODN * (a % MODN) % MODN
	}
}

func countGoodNumbers(n int64) int {
	even := 5
	odd := 4
	if n%2 == 0 {
		return int(powerMOD(uint64(even), uint64(n/2)) * powerMOD(uint64(odd), uint64(n/2)) % MODN)
	} else {
		return int(powerMOD(uint64(even), uint64(n/2+1)) * powerMOD(uint64(odd), uint64(n/2)) % MODN)
	}
}
func main() {
	fmt.Println("vim-go")
}
