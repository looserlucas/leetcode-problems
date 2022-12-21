package main

import (
	"container/heap"
	"fmt"
)

func less(a, b string) bool {
	if len(a) < len(b) {
		return true
	} else if len(a) > len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return false
		} else if a[i] < b[i] {
			return true
		}
	}
	return false
}

// An IntHeap is a min-heap of ints.
type IntHeap []string

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return less(h[i], h[j]) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(string))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func CreateHeap() *IntHeap {
	h := &IntHeap{}
	heap.Init(h)
	return h
}

var rank, parent []int

func find(i int) int {
	if parent[i] == i {
		return i
	} else {
		p := find(parent[i])
		parent[i] = p
		return p
	}
}

func union(i, j int) {
	ip := find(i)
	jp := find(j)
	if ip == jp {
		return
	}

	ri := rank[ip]
	rj := rank[jp]
	if ri < rj {
		parent[ip] = jp
	} else {
		parent[jp] = ip
		if ri == rj {
			rank[ip]++
		}
	}
}

func accountsMerge(a [][]string) [][]string {
	var emails []string
	index := make(map[string]int)
	firstSeen := make(map[string]int)
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a[i]); j++ {
			_, ok := firstSeen[a[i][j]]
			if !ok {
				firstSeen[a[i][j]] = i
				emails = append(emails, a[i][j])
				index[a[i][j]] = len(emails) - 1
			}
		}

	}
	rank = make([]int, len(emails))
	parent = make([]int, len(emails))
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	for i := 0; i < len(a); i++ {
		first := index[a[i][1]]
		for j := 2; j < len(a[i]); j++ {
			union(index[a[i][j]], first)
		}
	}

	type acc struct {
		name  string
		email *IntHeap
	}

	var res = make(map[int]acc)
	for i := 0; i < len(emails); i++ {
		p := find(i)
		_, ok := res[p]
		if !ok {
			h := CreateHeap()
			heap.Push(h, emails[i])
			res[p] = acc{name: a[firstSeen[emails[i]]][0], email: h}
		} else {
			heap.Push(res[p].email, emails[i])
		}
	}

	var ans = make([][]string, len(res))
	i := 0
	for _, v := range res {
		var tmp []string
		tmp = append(tmp, v.name)
		for v.email.Len() > 0 {
			tmp = append(tmp, heap.Pop(v.email).(string))
		}
		ans[i] = tmp
		i++
	}
	return ans
}

func main() {
	fmt.Println(accountsMerge([][]string{{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"}, {"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"}, {"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"}, {"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"}, {"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"}}))
	fmt.Println(accountsMerge([][]string{{"David", "David0@m.co", "David1@m.co"}, {"David", "David3@m.co", "David4@m.co"}, {"David", "David4@m.co", "David5@m.co"}, {"David", "David2@m.co", "David3@m.co"}, {"David", "David1@m.co", "David2@m.co"}}))
}
