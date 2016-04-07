package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/SlyMarbo/rss"
)

var firstRev = regexp.MustCompile(` rev -0+: `) // Matches the initial version of an RFC, usually " rev -00: "
var extractTitle = regexp.MustCompile(`".*"`)   // Extracts title of RFC from RSS title, eg "Title" - Author

type DraftRFCFilter struct {
	wchan chan string
}

// Filter for Draft RFCs that matches only the initial draft publication (not
// updates to an existing draft) and extracts just the title of the RFC
// excluding author's names.
func (f *DraftRFCFilter) Filter(item *rss.Item) {
	if firstRev.MatchString(item.Content) {
		log.Println("DraftRFCFilter first rev:", item.Title)
		title := extractTitle.FindString(item.Title)
		f.wchan <- fmt.Sprintf("New Draft: %s %s", title, item.Link)
	} else {
		log.Println("DraftRFCFilter not first rev:", item.Title)
	}
}

func ietfDraftRFC(wchan chan string) {
	ietf_rfc, err := NewRSS("http://tools.ietf.org/html/new-ids.rss")
	if err != nil {
		log.Fatalln("Error initialising IETF Draft RFC:", err)
	}
	f := &DraftRFCFilter{wchan: wchan}
	ietf_rfc.Read(f)
}
