package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/SlyMarbo/rss"
)

var firstRev = regexp.MustCompile(` rev -0+: `) // Matches the initial version of an RFC, usually " rev -00: "
var extractTitle = regexp.MustCompile(`".*"`)   // Extracts title of RFC from RSS title, eg "Title" - Author

type DraftRFCFilter struct{}

// Filter for Draft RFCs that matches only the initial draft publication (not
// updates to an existing draft) and extracts just the title of the RFC
// excluding author's names.
func (f *DraftRFCFilter) Filter(item *rss.Item) string {
	if firstRev.MatchString(item.Content) {
		title := extractTitle.FindString(item.Title)
		return fmt.Sprintf("%s %s", title, item.Link)
	}
	return ""
}

func ietfDraftRFC() <-chan string {
	ietf_rfc, err := NewRSS("http://tools.ietf.org/html/new-ids.rss")
	if err != nil {
		log.Fatalln("Error initialising IETF Draft RFC:", err)
	}
	go ietf_rfc.Read(&DraftRFCFilter{})

	return ietf_rfc.rchan
}
