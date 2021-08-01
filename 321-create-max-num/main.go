package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}
type Stack struct {
	Head   *Node
	Length int
}

func (s *Stack) Push(val int) {
	newNode := &Node{
		Value: val,
	}

	if s.Head == nil {
		s.Head = newNode
	} else {
		tmp := s.Head
		s.Head = newNode
		newNode.Next = tmp
	}
	s.Length++
}

func (s *Stack) Pop() int {
	if s.Head == nil {
		return -1
	}
	tmp := s.Head
	s.Head = tmp.Next
	s.Length--
	return tmp.Value
}

func CreateNewStack() (s *Stack) {
	return &Stack{
		Length: 0,
	}
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	var nums1Comp [2000][]int
	nums1Comp[0] = nil
	for i := 1; i <= k; i++ {
		var tmp []int
		if i > len(nums1) {
			nums1Comp[i] = nil
			break
		}
		st1 := CreateNewStack()
		for j := 0; j < len(nums1); j++ {
			if st1.Head == nil {
				st1.Push(nums1[j])
				continue
			}
			for st1.Head != nil && nums1[j] > st1.Head.Value && (st1.Length+len(nums1)-j-1 >= i) {
				st1.Pop()
			}
			st1.Push(nums1[j])
		}
		now := st1.Head
		for now != nil {
			tmp = append(tmp, now.Value)
			now = now.Next
		}
		var tmp1 []int
		for j := len(tmp) - 1; j >= len(tmp)-i; j-- {
			tmp1 = append(tmp1, tmp[j])
		}

		nums1Comp[i] = tmp1
	}

	var nums2Comp [2000][]int
	nums2Comp[0] = nil
	for i := 1; i <= k; i++ {
		if i > len(nums2) {
			nums2Comp[i] = nil
			break
		}
		st1 := CreateNewStack()
		for j := 0; j < len(nums2); j++ {
			if st1.Head == nil {
				st1.Push(nums2[j])
				continue
			}
			for st1.Head != nil && nums2[j] > st1.Head.Value && (st1.Length+len(nums2)-j-1 >= i) {
				st1.Pop()
			}
			st1.Push(nums2[j])
		}

		var tmp []int
		now := st1.Head
		for now != nil {
			tmp = append(tmp, now.Value)
			now = now.Next
		}
		var tmp1 []int
		for j := len(tmp) - 1; j >= len(tmp)-i; j-- {
			tmp1 = append(tmp1, tmp[j])
		}

		nums2Comp[i] = tmp1
	}

	var res []int
	// now merge
	for i := 0; i <= k; i++ {
		tmp := MergeArray(nums1Comp[i], nums2Comp[k-i])
		res = CompareArr(res, tmp)
	}
	return res
}

func CompareArr(a, b []int) []int {
	if len(a) > len(b) {
		return a
	}
	if len(b) > len(a) {
		return b
	}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return a
		}
		if a[i] < b[i] {
			return b
		}
	}
	return a
}

func MergeArray(a, b []int) []int {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	pt1 := 0
	pt2 := 0
	var res []int

	for pt1 < len(a) || pt2 < len(b) {
		fmt.Println(pt1)
		fmt.Println(pt2)
		fmt.Println(a[pt1])
		fmt.Println(b[pt2])
		if pt1 >= len(a) {
			res = append(res, b[pt2])
			pt2++
			continue
		}
		if pt2 >= len(b) {
			res = append(res, a[pt1])
			pt1++
			continue
		}
		if a[pt1] > b[pt2] {
			res = append(res, a[pt1])
			pt1++
		} else {
			if a[pt1] == b[pt2] {
				nx1 := 0
				nx2 := 0
				pt11 := pt1
				pt21 := pt2
				ok1 := false
				ok2 := false
				for (nx1 == nx2) || (pt11 < len(a) || pt21 < len(b)) {
					if pt11 >= len(a)-1 {
						nx1 = -1
						if !ok1 {
							pt11++
							ok1 = true
						}
					} else {
						nx1 = a[pt11+1]
						pt11++
					}
					if pt21 >= len(b)-1 {
						nx2 = -1
						if !ok2 {
							pt21++
							ok2 = true
						}
					} else {
						nx2 = b[pt21+1]
						pt21++
					}
				}

				fmt.Println("last one here")
				if nx1 > nx2 {
					for pt1 < pt11 {
						res = append(res, a[pt1])
						pt1++
					}
					continue
				} else {
					for pt2 < pt21 {
						res = append(res, b[pt2])
						pt2++
					}
					continue
				}
			}
			res = append(res, b[pt2])
			pt2++
		}
		fmt.Println(res)
	}

	fmt.Println("out of for loop")
	return res
}

/*
func MergeArray(a, b []int) []int {
	var ans []int
	for len(a) > 0 || len(b) > 0 {
		if a[0] > b[0] {
			ans = append(ans, a[0])
			a = a[1:]
		} else {
			ans = append(ans, b[0])
			b = b[1:]
		}
	}
	return ans
}
*/
func main() {
	/*
		nums1 := []int{9, 1, 2, 5, 8, 3}
		nums2 := []int{3, 4, 6, 5}
		k := 5
	*/
	/*
						nums1 := []int{2, 5, 6, 4, 4, 0}
						nums2 := []int{7, 3, 8, 0, 6, 5, 7, 6, 2}
						k := 15
				[2,1,7,8,0,1,7,3,5,8,9,0,0,7,0,2,2,7,3,5,5]
				[2,6,2,0,1,0,5,4,5,5,3,3,3,4]
				35
		[2,0,2,1,2,2,2,2,0,1,0,0,2,0,2,0,2,1,0,1,1,0,1,0,1,2,1,1,1,0,1,2,2,1,0,0,1,2,1,2,2,1,1,0,1,2,0,2,0,1,2,0,2,1,1,1,2,0,0,1,0,2,1,2,0,1,0,0,0,1,2,1,0,1,1,2,0,2,2,0,0,1,1,2,2,1,1,2,2,1,0,1,2,0,1,2,2,0,0,0,2,0,2,0,2,2,0,1,1,1,1,2,2,2,2,0,0,2,2,2,2,0,2,0,1,0,0,2,1,0,0,2,0,2,1,1,1,1,0,1,2,0,2,1,0,1,1,1,0,0,2,2,2,0,2,1,1,1,2,2,0,0,2,2,2,2,2,0,2,0,2,0,2,0,0,1,0,1,1,0,0,2,1,1,2,2,2,1,2,2,0,0,2,1,0,2,1,2,1,1,1,0,2,0,1,1,2,1,1,0,0,1,0,1,2,2,2,0,2,2,1,0,1,2,1,2,0,2,2,0,1,2,2,1,2,2,1,1,2,2,2,2,2,1,2,0,1,1,1,2,2,2,0,2,0,2,0,2,1,1,0,2,2,2,1,0,2,1,2,2,2,0,1,1,1,1,1,1,0,0,0,2,2,0,1,2,1,0,0,2,2,2,2,1,0,2,0,1,2,0]
		[1,1,1,0,0,1,1,0,2,1,0,1,2,1,0,2,2,1,0,2,0,1,1,0,0,2,2,0,1,0,2,0,2,2,2,2,1,1,1,1,0,0,0,0,2,1,0,2,1,1,2,1,2,2,0,2,1,0,2,0,0,2,0,2,2,1,0,1,0,0,2,1,1,1,2,2,0,0,0,1,1,2,0,2,2,0,1,0,2,1,0,2,1,1,1,0,1,1,2,0,2,0,1,1,2,0,2,0,1,2,1,0,2,0,1,0,0,0,1,2,1,2,0,1,2,2,1,1,0,1,2,1,0,0,1,0,2,2,1,2,2,0,0,0,2,0,0,0,1,0,2,0,2,1,0,0,1,2,0,1,1,0,1,0,2,2,2,1,1,0,1,1,2,1,0,2,2,2,1,2,2,2,2,0,1,1,0,1,2,1,2,2,0,0,0,0,0,1,1,1,2,1,2,1,1,0,1,2,0,1,2,1,2,2,2,2,0,0,0,0,2,0,1,2,0,1,1,1,1,0,1,2,2,1,0,1,2,2,1,2,2,2,0,2,0,1,1,2,0,0,2,2,0,1,0,2,1,0,0,1,1,1,1,0,0,2,2,2,2,0,0,1,2,1,1,2,0,1,2,1,0,2,0,0,2,1,1,0,2,1,1,2,2,0,1,0,2,0,1,0]
		600
	*/
	/*
		nums1 := []int{2, 1, 7, 8, 0, 1, 7, 3, 5, 8, 9, 0, 0, 7, 0, 2, 2, 7, 3, 5, 5}
		nums2 := []int{2, 6, 2, 0, 1, 0, 5, 4, 5, 5, 3, 3, 3, 4}
		res := []int{2, 6, 2, 2, 1, 7, 8, 0, 1, 7, 3, 5, 8, 9, 0, 1, 0, 5, 4, 5, 5, 3, 3, 3, 4, 0, 0, 7, 0, 2, 2, 7, 3, 5, 5}
		k := 35
	*/
	nums1 := []int{2, 0, 2, 1, 2, 2, 2, 2, 0, 1, 0, 0, 2, 0, 2, 0, 2, 1, 0, 1, 1, 0, 1, 0, 1, 2, 1, 1, 1, 0, 1, 2, 2, 1, 0, 0, 1, 2, 1, 2, 2, 1, 1, 0, 1, 2, 0, 2, 0, 1, 2, 0, 2, 1, 1, 1, 2, 0, 0, 1, 0, 2, 1, 2, 0, 1, 0, 0, 0, 1, 2, 1, 0, 1, 1, 2, 0, 2, 2, 0, 0, 1, 1, 2, 2, 1, 1, 2, 2, 1, 0, 1, 2, 0, 1, 2, 2, 0, 0, 0, 2, 0, 2, 0, 2, 2, 0, 1, 1, 1, 1, 2, 2, 2, 2, 0, 0, 2, 2, 2, 2, 0, 2, 0, 1, 0, 0, 2, 1, 0, 0, 2, 0, 2, 1, 1, 1, 1, 0, 1, 2, 0, 2, 1, 0, 1, 1, 1, 0, 0, 2, 2, 2, 0, 2, 1, 1, 1, 2, 2, 0, 0, 2, 2, 2, 2, 2, 0, 2, 0, 2, 0, 2, 0, 0, 1, 0, 1, 1, 0, 0, 2, 1, 1, 2, 2, 2, 1, 2, 2, 0, 0, 2, 1, 0, 2, 1, 2, 1, 1, 1, 0, 2, 0, 1, 1, 2, 1, 1, 0, 0, 1, 0, 1, 2, 2, 2, 0, 2, 2, 1, 0, 1, 2, 1, 2, 0, 2, 2, 0, 1, 2, 2, 1, 2, 2, 1, 1, 2, 2, 2, 2, 2, 1, 2, 0, 1, 1, 1, 2, 2, 2, 0, 2, 0, 2, 0, 2, 1, 1, 0, 2, 2, 2, 1, 0, 2, 1, 2, 2, 2, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 2, 2, 0, 1, 2, 1, 0, 0, 2, 2, 2, 2, 1, 0, 2, 0, 1, 2, 0}
	nums2 := []int{
		1, 1, 1, 0, 0, 1, 1, 0, 2, 1, 0, 1, 2, 1, 0, 2, 2, 1, 0, 2, 0, 1, 1, 0, 0, 2, 2, 0, 1, 0, 2, 0, 2, 2, 2, 2, 1, 1, 1, 1, 0, 0, 0, 0, 2, 1, 0, 2, 1, 1, 2, 1, 2, 2, 0, 2, 1, 0, 2, 0, 0, 2, 0, 2, 2, 1, 0, 1, 0, 0, 2, 1, 1, 1, 2, 2, 0, 0, 0, 1, 1, 2, 0, 2, 2, 0, 1, 0, 2, 1, 0, 2, 1, 1, 1, 0, 1, 1, 2, 0, 2, 0, 1, 1, 2, 0, 2, 0, 1, 2, 1, 0, 2, 0, 1, 0, 0, 0, 1, 2, 1, 2, 0, 1, 2, 2, 1, 1, 0, 1, 2, 1, 0, 0, 1, 0, 2, 2, 1, 2, 2, 0, 0, 0, 2, 0, 0, 0, 1, 0, 2, 0, 2, 1, 0, 0, 1, 2, 0, 1, 1, 0, 1, 0, 2, 2, 2, 1, 1, 0, 1, 1, 2, 1, 0, 2, 2, 2, 1, 2, 2, 2, 2, 0, 1, 1, 0, 1, 2, 1, 2, 2, 0, 0, 0, 0, 0, 1, 1, 1, 2, 1, 2, 1, 1, 0, 1, 2, 0, 1, 2, 1, 2, 2, 2, 2, 0, 0, 0, 0, 2, 0, 1, 2, 0, 1, 1, 1, 1, 0, 1, 2, 2, 1, 0, 1, 2, 2, 1, 2, 2, 2, 0, 2, 0, 1, 1, 2, 0, 0, 2, 2, 0, 1, 0, 2, 1, 0, 0, 1, 1, 1, 1, 0, 0, 2, 2, 2, 2, 0, 0, 1, 2, 1, 1, 2, 0, 1, 2, 1, 0, 2, 0, 0, 2, 1, 1, 0, 2, 1, 1, 2, 2, 0, 1, 0, 2, 0, 1, 0}
	k := 600
	fmt.Println(maxNumber(nums1, nums2, k))
	//fmt.Println(res)
}
