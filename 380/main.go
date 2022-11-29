package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	index map[int]int
	list  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		index: make(map[int]int),
		list:  make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.index[val]
	if ok {
		return false
	}
	this.list = append(this.list, val)
	this.index[val] = len(this.list) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	v, ok := this.index[val]
	if !ok {
		return false
	}
	// swat last element of the list and the soon-to-be-removed element places
	this.index[this.list[len(this.list)-1]] = v
	this.list[v], this.list[len(this.list)-1] = this.list[len(this.list)-1], this.list[v]
	this.list = this.list[:len(this.list)-1]
	delete(this.index, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.list))
	return this.list[i]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
func main() {
	fmt.Println("vim-go")
}
