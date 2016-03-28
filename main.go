package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	log.Println("Starting RFC-Bot")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	log.Println("Initialising Writers")
	writer := NewWriter()

	log.Println("Initialising Twitter")
	twitter, err := NewTwitter(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"), os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	if err != nil {
		log.Fatal("Could not authenticate to twitter:", err)
	}
	writer.AddWriter(twitter)

	log.Println("Initialising Readers")
	go ietfRFC(writer.Wchan)
	go ietfDraftRFC(writer.Wchan)

	log.Println("Initialised")
	for {
	}
}
