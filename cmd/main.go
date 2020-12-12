// Package main -
package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/shanehowearth/go-pls-announce-bot/publish"
	"github.com/shanehowearth/go-pls-announce-bot/publish/twitter"
	"github.com/shanehowearth/go-pls-announce-bot/storage"
)

func main() {
	// Collect twitter credentials
	twitterAPIKey := os.Getenv("TWITTER_API_KEY")
	twitterAPISecret := os.Getenv("TWITTER_API_SECRET")
	twitterAccessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	twitterAccessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	// Create the twitter client
	tc, err := twitter.NewTwitterClient(twitterAPIKey, twitterAPISecret, twitterAccessToken, twitterAccessTokenSecret)
	if err != nil {
		log.Fatalf("Unable to create twitter client with error %v", err)
	}
	platforms := []publish.Publisher{}
	platforms = append(platforms, tc)

	// Connect to the data store
	pgDB := os.Getenv("POSTGRES_DATABASE")
	db, err := postgresstore.NewPGStore(pgDB)
	if err != nil {
		log.Fatalf("Unable to create postgres connection with error %v", err)
	}

	// Retrieve content
	set := flag.String("type", "", "the name of the content set to produce")
	flag.Parse()
	data, translation, explanation, err := storage.GetContent(*set, db)
	if err != nil {
		log.Fatalf("Unable to retrieve content with error %v", err)
	}

	content := strings.Join([]string{data, translation, explanation}, "\n\n")
	// Publish content
	for idx := range platforms {
		err := platforms[idx].PublishContent(content)
		if err != nil {
			// not the end of the world so don't die, but someone somewhere
			// should have a look at it
			log.Printf("ERROR: Getting content for %v returned error %v", platforms[idx], err)
		}
	}

}
