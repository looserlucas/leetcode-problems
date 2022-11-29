package main

import "fmt"

type DisjoinSet struct {
	// parent[i] is the parent of node i
	parent []int
	// rank[i] is the rank of the tree with root i
	rank []int
}

func (d *DisjoinSet) find(i int) int {
	if d.parent[i] == i {
		return i
	} else {
		res := d.find(d.parent[i])
		// compress parent
		d.parent[i] = res
		return res
	}
}

func (d *DisjoinSet) union(i, j int) {
	irep := d.find(i)
	jrep := d.find(j)
	if irep == jrep {
		return
	}

	irank := d.rank[irep]
	jrank := d.rank[jrep]

	// merge smaller rank tree -> bigger rank tree
	// if irank == jrank, merge 2 tree and increase the rank of the new tree by 1
	if irank < jrank {
		d.parent[irep] = jrep
	} else if jrank < irank {
		d.parent[jrep] = irep
	} else {
		d.parent[irep] = jrep
		d.rank[jrep]++
	}
}

func CreateDisjoinSet(length int) *DisjoinSet {
	parent := make([]int, length)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	rank := make([]int, length)
	return &DisjoinSet{
		parent: parent,
		rank:   rank,
	}
}

func largestComponentSize(nums []int) int {
	ds := CreateDisjoinSet(100001)
	for i := 0; i < len(nums); i++ {
		ds.union(nums[i], nums[i])
		for j := 2; j*j <= nums[i]; j++ {
			if nums[i]%j == 0 {
				if j*j == nums[i] {
					ds.union(j, nums[i])
				} else {
					ds.union(j, nums[i])
					ds.union(nums[i]/j, nums[i])
				}
			}
		}
	}
	var count = make(map[int]int)
	var res = 1
	for i := 0; i < len(nums); i++ {
		p := ds.find(nums[i])
		count[p]++
		if count[p] > res {
			res = count[p]
		}
	}
	return res
}

func main() {
	fmt.Println(largestComponentSize([]int{4, 6, 15, 35}))
	fmt.Println(largestComponentSize([]int{20, 50, 9, 63}))
	fmt.Println(largestComponentSize([]int{2, 3, 6, 7, 4, 12, 21, 39}))
}
