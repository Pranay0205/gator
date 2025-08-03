package main

import (
	"context"
	"fmt"

	"github.com/Pranay0205/gator/internal/rss"
)

const feedURL = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := rss.FetchFeed(context.Background(), feedURL)

	if err != nil {
		return fmt.Errorf("failed to fetch RSS feed: %w", err)
	}

	printRSSFeed(*rssFeed)

	return nil

}

func printRSSFeed(rssFeed rss.RSSFeed){
	fmt.Printf("Title: %s\n", rssFeed.Channel.Title)
	fmt.Printf("Description: %s\n", rssFeed.Channel.Description)
	for i := range rssFeed.Channel.Item {
		fmt.Printf("\t- Title: %s\n", rssFeed.Channel.Item[i].Title)
		fmt.Printf("\t- Description: %s\n", rssFeed.Channel.Item[i].Description)
		fmt.Printf("\t- Link: %s\n", rssFeed.Channel.Item[i].Link)
		fmt.Printf("\t- Publish Date: %s\n", rssFeed.Channel.Item[i].PubDate)
	}
}