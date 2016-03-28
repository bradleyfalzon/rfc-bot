package main

import (
	"fmt"
	"log"
	"time"

	"github.com/SlyMarbo/rss"
)

type rssFeed struct {
	feed  *rss.Feed
	rchan chan string
}

func NewRSS(url string) (*rssFeed, error) {
	rssFeed := &rssFeed{}
	var err error
	rssFeed.feed, err = rss.Fetch(url)
	if err != nil {
		return nil, err
	}
	rssFeed.rchan = make(chan string)
	rssFeed.MarkAllRead()
	return rssFeed, nil
}

// MarkAllRead marks all items in feed as read and decrements unread counter
func (r *rssFeed) MarkAllRead() {
	for k, _ := range r.feed.Items {
		r.feed.Items[k].Read = true
	}
	r.feed.Unread = 0
}

// Read periodically refreshes rss feed looking for new items
//func (r *rssFeed) Read(wchan chan<- string) {
func (r *rssFeed) Read() {
	for {
		// Sleep until the next refresh period
		sleep := r.feed.Refresh.Sub(time.Now())
		log.Printf("%s next update at: %s (%s)", r.feed.Title, r.feed.Refresh, sleep)
		time.Sleep(sleep)

		r.feed.Update()
		for _, item := range r.feed.Items {
			if !item.Read {
				item.Read = true
				r.feed.Unread--
				r.rchan <- fmt.Sprintf("%s %s", item.Title, item.Link)
			}
		}
	}
}
