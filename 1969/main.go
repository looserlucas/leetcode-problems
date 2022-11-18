package main

import "fmt"

const MODN = 1e9 + 7

func modPow(x, y uint64, m uint64) uint64 {
	if y == 0 {
		return 1
	}
	if y == 1 {
		return x % m
	}
	p := modPow(x, y/2, m)
	p = (p * p) % m
	if y%2 == 1 {
		return p * (x % m) % m
	} else {
		return p
	}
}

// res = (2^n-2)^(2^(n-1) -1) * (2^n-1)
func minNonZeroProduct(p int) int {
	var a, b, c uint64 = 1, 0, 0
	for i := 1; i <= p; i++ {
		a *= 2
	}
	a = a - 1
	b = a / 2
	c = a - 1
	fmt.Println(a)
	return (int(modPow(c, b, MODN)) * int(a%MODN)) % MODN
}

func main() {
	fmt.Println(minNonZeroProduct(35))
	fmt.Println(minNonZeroProduct(60))
}
