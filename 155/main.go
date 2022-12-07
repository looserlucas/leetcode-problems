package main

import "fmt"

type MinStack struct {
	cur []int
	min []int
}

func Constructor() MinStack {
	return MinStack{
		cur: make([]int, 0),
		min: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.cur = append(this.cur, val)
	n := len(this.cur)
	m := len(this.min)
	if m == 0 || this.cur[n-1] < this.cur[this.min[m-1]] {
		this.min = append(this.min, n-1)
	}
}

func (this *MinStack) Pop() {
	n := len(this.cur)
	m := len(this.min)
	if this.min[m-1] == n-1 {
		this.min = this.min[0 : m-1]
	}
	this.cur = this.cur[0 : n-1]
}

func (this *MinStack) Top() int {
	return this.cur[len(this.cur)-1]
}

func (this *MinStack) GetMin() int {
	return this.cur[this.min[len(this.min)-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	fmt.Print(s.GetMin())
	s.Pop()
	fmt.Print(s.Top())
	fmt.Print(s.GetMin())
}
