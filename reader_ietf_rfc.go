package main

import "log"

func ietfRFC() <-chan string {
	ietf_rfc, err := NewRSS("http://tools.ietf.org/html/new-rfcs.rss")
	if err != nil {
		log.Fatalln("Error initialising IETF RFC:", err)
	}
	go ietf_rfc.Read()
	return ietf_rfc.rchan
}
