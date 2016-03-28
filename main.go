package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	twitter, err := NewTwitter(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"), os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	if err != nil {
		log.Fatal("Could not authenticate to twitter:", err)
	}

	writer := NewWriter()
	writer.AddWriter(twitter)

	for {
		select {
		case str := <-ietfRFC():
			writer.Wchan <- str
		case str := <-ietfDraftRFC():
			writer.Wchan <- str
		}
	}
}
