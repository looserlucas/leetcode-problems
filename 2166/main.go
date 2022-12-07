package main

import (
	"fmt"
	"strings"
)

type Bitset struct {
	count1  int
	set     []bool
	flipped bool
}

func Constructor(size int) Bitset {
	return Bitset{
		count1:  0,
		flipped: false,
		set:     make([]bool, size),
	}
}

func (this *Bitset) Fix(idx int) {
	if this.set[idx] == this.flipped {
		this.set[idx] = !this.set[idx]
		this.count1++
	}
}

func (this *Bitset) Unfix(idx int) {
	if this.set[idx] != this.flipped {
		this.set[idx] = !this.set[idx]
		this.count1--
	}
}

func (this *Bitset) Flip() {
	this.flipped = !this.flipped
	this.count1 = len(this.set) - this.count1
}

func (this *Bitset) All() bool {
	return (this.count1 == len(this.set))
}

func (this *Bitset) One() bool {
	return (this.count1 > 0)
}

func (this *Bitset) Count() int {
	return this.count1
}

func (this *Bitset) ToString() string {
	var res strings.Builder
	for i := 0; i < len(this.set); i++ {
		if this.set[i] != this.flipped {
			res.WriteRune('1')
		} else {
			res.WriteRune('0')
		}
	}
	return res.String()
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */
func main() {
	fmt.Println("vim-go")
}
