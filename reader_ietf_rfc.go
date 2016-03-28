package main

import (
	"fmt"
	"log"

	"github.com/SlyMarbo/rss"
)

type IETFNewFilter struct{}

func (f *IETFNewFilter) Filter(item *rss.Item) string {
	return fmt.Sprintf("RFC: %s %s", item.Title, item.Link)
}

func ietfRFC() <-chan string {
	ietf_rfc, err := NewRSS("http://tools.ietf.org/html/new-rfcs.rss")
	if err != nil {
		log.Fatalln("Error initialising IETF RFC:", err)
	}
	go ietf_rfc.Read(&IETFNewFilter{})
	return ietf_rfc.rchan
}
