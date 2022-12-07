package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []*Tweet

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].Timestamp > h[j].Timestamp }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Tweet))
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

type Twitter struct {
	CurTimeStamp int
	Users        map[int]*User
}

type User struct {
	Id        int
	Following map[int]*User
	TweetList *Tweet
}

type Tweet struct {
	Id        int
	Timestamp int
	Next      *Tweet
}

func Constructor() Twitter {
	users := make(map[int]*User)
	return Twitter{
		Users:        users,
		CurTimeStamp: 0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.Users[userId]; !ok {
		newUser := &User{
			Id:        userId,
			Following: make(map[int]*User),
			TweetList: nil,
		}
		this.Users[userId] = newUser
	}
	newTweet := &Tweet{
		Id:        tweetId,
		Timestamp: this.CurTimeStamp,
		Next:      this.Users[userId].TweetList,
	}
	this.Users[userId].TweetList = newTweet

	this.CurTimeStamp++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	user, ok := this.Users[userId]
	if !ok {
		return nil
	}

	h := CreateHeap()
	if user.TweetList != nil {
		heap.Push(h, user.TweetList)
	}
	for _, v := range user.Following {
		if v.TweetList != nil {
			heap.Push(h, v.TweetList)
		}
	}
	var res []int
	for i := 0; i < 10; i++ {
		if h.Len() > 0 {
			x := heap.Pop(h).(*Tweet)
			if x.Next != nil {
				heap.Push(h, x.Next)
			}
			res = append(res, x.Id)
		}
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.Users[followerId]; !ok {
		newUser := &User{
			Id:        followerId,
			Following: make(map[int]*User),
			TweetList: nil,
		}
		this.Users[followerId] = newUser
	}
	if _, ok := this.Users[followeeId]; !ok {
		newUser := &User{
			Id:        followeeId,
			Following: make(map[int]*User),
			TweetList: nil,
		}
		this.Users[followeeId] = newUser
	}
	this.Users[followerId].Following[followeeId] = this.Users[followeeId]
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.Users[followerId]; !ok {
		return
	}
	if _, ok := this.Users[followeeId]; !ok {
		return
	}

	delete(this.Users[followerId].Following, followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
func main() {
	fmt.Println("vim-go")
}
