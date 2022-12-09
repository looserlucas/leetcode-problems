package main

import (
	"fmt"
)

type LUPrefix struct {
	n      int
	parent []int
	total  []int
}

func (this *LUPrefix) find(i int) int {
	if this.parent[i] == i {
		return i
	} else {
		p := this.find(this.parent[i])
		this.parent[i] = p
		return p
	}
}

func (this *LUPrefix) union(i, j int) {
	ip := this.find(i)
	jp := this.find(j)
	if ip == jp {
		return
	}

	if ip > jp {
		this.parent[ip] = jp
		this.total[jp] += this.total[ip]
	} else {
		this.parent[jp] = ip
		this.total[ip] += this.total[jp]
	}
}

func Constructor(n int) LUPrefix {
	parent := make([]int, n+1)
	for i := 1; i < len(parent); i++ {
		parent[i] = i
	}
	total := make([]int, n+1)
	return LUPrefix{
		n:      n,
		parent: parent,
		total:  total,
	}
}

func (this *LUPrefix) Upload(video int) {
	this.total[video] = 1
	if video-1 > 0 && this.total[video-1] > 0 {
		this.union(video-1, video)
	}
	if video+1 <= this.n && this.total[video+1] > 0 {
		this.union(video, video+1)
	}
	fmt.Println(video, this.parent)
	fmt.Println(video, this.total)
}

func (this *LUPrefix) Longest() int {
	return this.total[1]
}

/**
 * Your LUPrefix object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.Upload(video);
 * param_2 := obj.Longest();
 */
func main() {
	l := Constructor(4)
	l.Upload(3)
	fmt.Println(l.Longest())
	l.Upload(1)
	fmt.Println(l.Longest())
	l.Upload(2)
	fmt.Println(l.Longest())
}
