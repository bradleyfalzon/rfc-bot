package main

import (
	"fmt"
	"log"

	"github.com/SlyMarbo/rss"
)

type IETFNewFilter struct {
	wchan chan string
}

func (f *IETFNewFilter) Filter(item *rss.Item) {
	f.wchan <- fmt.Sprintf("New RFC: %s %s", item.Title, item.Link)
}

func ietfRFC(wchan chan string) {
	ietf_rfc, err := NewRSS("http://tools.ietf.org/html/new-rfcs.rss")
	if err != nil {
		log.Fatalln("Error initialising IETF RFC:", err)
	}
	f := &IETFNewFilter{wchan: wchan}
	ietf_rfc.Read(f)
}
