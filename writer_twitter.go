package main

import (
	"log"
	"net/url"
	"regexp"

	"github.com/ChimeraCoder/anaconda"
)

type Twitter struct {
	api *anaconda.TwitterApi
}

func NewTwitter(consumerKey, consumerSecret, accessToken, accessTokenSecret string) (*Twitter, error) {
	t := &Twitter{}
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	t.api = anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	_, err := t.api.VerifyCredentials()
	return t, err
}

func (t Twitter) Write(str []byte) (int, error) {
	log.Printf("Before hashing: %q", string(str))
	tweet := prependHashes(string(str))
	log.Printf("Tweeting: %q", tweet)

	t.api.PostTweet(tweet, url.Values{})
	return len(str), nil
}

// Begins with uppercase letter, then many upper-lower-numeric-slash
var acronym = regexp.MustCompile(`\b([A-Z][a-zA-Z0-9/]*[A-Z0-9]+)`)

func prependHashes(str string) string {
	return acronym.ReplaceAllString(str, "#$1")
}
