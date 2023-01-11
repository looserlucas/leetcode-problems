package main

import "fmt"

func minimumBuckets(hamsters string) int {
	var a []byte = make([]byte, len(hamsters))
	for i := 0; i < len(hamsters); i++ {
		a[i] = hamsters[i]
	}
	var res int
	for i := 0; i < len(a); i++ {
		if a[i] == 'H' {
			if i-1 >= 0 && a[i-1] == 'F' {
				continue
			}
			if i+1 < len(a) && a[i+1] == '.' {
				a[i+1] = 'F'
				res++
				continue
			}
			if i-1 >= 0 && a[i-1] == '.' && ((i+1 < len(a) && a[i+1] == 'H') || (i+1 == len(a))) {
				a[i-1] = 'F'
				res++
				continue
			}
			return -1
		}
	}
	return res
}

func main() {
	fmt.Println(minimumBuckets(".H.H."))
	fmt.Println(minimumBuckets("H..H"))
	fmt.Println(minimumBuckets(".HHH."))
}
