package main

import "fmt"

type TweetCounts struct {
	tweet map[string][]int
}

func Constructor() TweetCounts {
	return TweetCounts{
		tweet: make(map[string][]int),
	}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	this.tweet[tweetName] = append(this.tweet[tweetName], time)
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	interval := 60
	if freq == "hour" {
		interval *= 60
	} else if freq == "day" {
		interval *= 60 * 24
	}
	var res []int
	for i := 0; i <= (endTime-startTime)/interval; i++ {
		res = append(res, 0)
	}

	tweets := this.tweet[tweetName]
	for _, tw := range tweets {
		if (startTime <= tw) && (tw <= endTime) {
			idx := (tw - startTime) / interval
			fmt.Println(idx, tw)
			res[idx]++
		}
	}
	return res
}

func main() {
	t := Constructor()
	t.RecordTweet("tweet3", 0)
	t.RecordTweet("tweet3", 60)
	t.RecordTweet("tweet3", 10)
	fmt.Println(t)
	//fmt.Println(t.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59))
	fmt.Println(t.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60))
}
