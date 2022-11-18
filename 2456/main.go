package main

import "fmt"

func greater(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return true
	} else if len(s1) < len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] > s2[i] {
			return true
		} else if s1[i] < s2[i] {
			return false
		}
	}
	return false
}
func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	type Metadata struct {
		CurrentTopVideoIndex int
		TotalViews           int
	}
	var listedUser = make(map[string]bool)
	var res [][]string
	var max = -1
	user := make(map[string]Metadata)
	for i := 0; i < len(creators); i++ {
		v, ok := user[creators[i]]
		if !ok {
			tmp := Metadata{
				CurrentTopVideoIndex: i,
				TotalViews:           views[i],
			}
			user[creators[i]] = tmp
		} else {
			if views[i] > views[v.CurrentTopVideoIndex] {
				tmp := Metadata{
					CurrentTopVideoIndex: i,
					TotalViews:           v.TotalViews + views[i],
				}
				user[creators[i]] = tmp
			} else if views[i] == views[v.CurrentTopVideoIndex] {
				if greater(ids[v.CurrentTopVideoIndex], ids[i]) {
					tmp := Metadata{
						CurrentTopVideoIndex: i,
						TotalViews:           v.TotalViews + views[i],
					}
					user[creators[i]] = tmp
				} else {
					tmp := Metadata{
						CurrentTopVideoIndex: v.CurrentTopVideoIndex,
						TotalViews:           v.TotalViews + views[i],
					}
					user[creators[i]] = tmp
				}
			} else {
				tmp := Metadata{
					CurrentTopVideoIndex: v.CurrentTopVideoIndex,
					TotalViews:           v.TotalViews + views[i],
				}
				user[creators[i]] = tmp
			}
		}
		v1 := user[creators[i]]
		if v1.TotalViews > max {
			max = v1.TotalViews
			listedUser = nil
			listedUser = make(map[string]bool)
			listedUser[creators[i]] = true
			res = nil
			res = append(res, []string{creators[i], ids[v1.CurrentTopVideoIndex]})
		} else if v1.TotalViews == max {
			if _, ok := listedUser[creators[i]]; !ok {
				res = append(res, []string{creators[i], ids[v1.CurrentTopVideoIndex]})
				listedUser[creators[i]] = true
			}
		}
	}
	return res
}

func main() {
	fmt.Println(mostPopularCreator([]string{"alice", "bob", "alice", "chris"}, []string{"one", "two", "three", "four"}, []int{5, 10, 5, 4}))
	fmt.Println(mostPopularCreator([]string{"alice", "alice", "alice"}, []string{"a", "b", "c"}, []int{1, 2, 2}))
	fmt.Println(mostPopularCreator([]string{"bb", "bb"}, []string{"baa", "bba"}, []int{1, 0}))
}
