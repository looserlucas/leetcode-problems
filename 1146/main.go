package main

import "sort"

type SnapshotArray struct {
	cur    []int
	snapId int
	snap   [][]int
	dirty  bool
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		cur:    make([]int, length),
		dirty:  true,
		snapId: -1,
		snap:   nil,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.dirty = true
	this.cur[index] = val
}

func (this *SnapshotArray) Snap() int {
	this.snapId++
	if this.dirty {
		newArr := make([]int, len(this.cur))
		copy(newArr, this.cur)
		newArr = append(newArr, this.snapId)
		this.snap = append(this.snap, newArr)
		this.dirty = false
	}
	return this.snapId
}

func (this *SnapshotArray) Get(index int, snapID int) int {
	x := sort.Search(len(this.snap), func(i int) bool {
		return this.snap[i][len(this.snap[i])-1] >= snapID
	})
	if x == len(this.snap) {
		return this.snap[len(this.snap)-1][index]
	} else {
		if this.snap[x][len(this.snap[x])-1] > snapID {
			return this.snap[x-1][index]
		} else {
			return this.snap[x][index]
		}
	}
}
