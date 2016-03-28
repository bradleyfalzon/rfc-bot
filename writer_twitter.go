package main

import (
	"net/url"

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
	t.api.PostTweet(string(str), url.Values{})
	return len(str), nil
}
