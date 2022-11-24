package main

type stack []int

func (s *stack) Pop() int {
	if len(*s) == 0 {
		return -1
	}

	x := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return x
}

func (s *stack) Push(x int) {
	(*s) = append((*s), x)
}

var id = 0
var res [][]int

var ids []int
var low []int
var onStack []bool
var s stack

// https://www.youtube.com/watch?v=TyWtx7q2D7Y
func Tarjan(n int, edges [][]int) {
	var adjV = make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		v, ok := adjV[edges[i][0]]
		if !ok {
			var tmp []int
			tmp = append(tmp, edges[i][1])
			adjV[edges[i][0]] = tmp
		} else {
			v = append(v, edges[i][1])
			adjV[edges[i][0]] = v
		}

		v, ok = adjV[edges[i][1]]
		if !ok {
			var tmp []int
			tmp = append(tmp, edges[i][0])
			adjV[edges[i][1]] = tmp
		} else {
			v = append(v, edges[i][0])
			adjV[edges[i][1]] = v
		}
	}

	ids = make([]int, n)
	low = make([]int, n)
	onStack = make([]bool, n)
	var tmp stack
	s = tmp
	for i := 0; i < len(ids); i++ {
		ids[i] = -1
	}

	for i := 0; i < n; i++ {
		if ids[i] == -1 {
			dfs(adjV, i)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func dfs(adjV map[int][]int, i int) {
	// house keeping step: push vertice i to stack, mark it as in the stack
	// update its id and low link value to the current visit id
	s.Push(i)
	onStack[i] = true
	ids[i] = id
	low[i] = id
	id++

	for _, v := range adjV[i] {
		if ids[v] == -1 {
			dfs(adjV, v)
		}
		// if v is onstack, which means i and v are in the same SCC
		// min the low link value of low[i] = min(low[i], low[v])
		if onStack[v] {
			low[i] = min(low[i], low[v])
		}
	}

	// check whether we are at the start of a SCC
	if ids[i] == low[i] {
		var tmpRes []int
		for len(s) > 0 {
			v := s.Pop()
			onStack[v] = false
			low[v] = ids[i]
			tmpRes = append(tmpRes, v)
			if v == i {
				break
			}
		}
		if tmpRes != nil {
			res = append(res, tmpRes)
		}
	}
}
