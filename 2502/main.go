package main

import "fmt"

type Allocator struct {
	mIDToMem map[int][]int
	mem      []bool
}

func Constructor(n int) Allocator {
	return Allocator{
		mIDToMem: make(map[int][]int),
		mem:      make([]bool, n),
	}
}

func (this *Allocator) Allocate(size int, mID int) int {
	for i := 0; i < len(this.mem); i++ {
		if !this.mem[i] && i+size-1 < len(this.mem) && !this.mem[i+size-1] {
			var ok bool = true
			for j := i; j < i+size; j++ {
				if this.mem[j] {
					ok = false
					break
				}
			}
			if ok {
				for j := i; j < i+size; j++ {
					this.mem[j] = true
				}
				this.mIDToMem[mID] = append(this.mIDToMem[mID], i)
				this.mIDToMem[mID] = append(this.mIDToMem[mID], size)
				return i
			}
		}
	}
	return -1
}

func (this *Allocator) Free(mID int) int {
	var res int
	for i := 0; i < len(this.mIDToMem[mID]); i += 2 {
		for j := this.mIDToMem[mID][i]; j < this.mIDToMem[mID][i]+this.mIDToMem[mID][i+1]; j++ {
			this.mem[j] = false
		}
		res += this.mIDToMem[mID][i+1]
	}
	delete(this.mIDToMem, mID)
	return res
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.Free(mID);
 */
func main() {
	m := Constructor(10)
	m.Allocate(1, 1)
	m.Allocate(1, 2)
	m.Allocate(1, 3)
	m.Free(2)
	fmt.Println(m.mIDToMem, m.mem)
	m.Allocate(3, 4)
	fmt.Println(m.mIDToMem, m.mem)
	m.Allocate(1, 1)
	m.Allocate(1, 1)
	m.Free(1)
	m.Allocate(10, 2)
	m.Free(7)
}
