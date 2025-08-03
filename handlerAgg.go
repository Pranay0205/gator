package main

import (
	"context"
	"fmt"

	"github.com/Pranay0205/gator/internal/rss"
)


func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	feedURL := cmd.Args[0]

	rssFeed, err := rss.FetchFeed(context.Background(), feedURL)

	if err != nil {
		return fmt.Errorf("failed to fetch RSS feed: %w", err)
	}

	printFeed(*rssFeed)

	return nil

}

func printFeed(rssFeed rss.RSSFeed){
	fmt.Printf("Title: %s\n", rssFeed.Channel.Title)
	fmt.Printf("Description: %s\n", rssFeed.Channel.Description)
	for i := range rssFeed.Channel.Item {
		fmt.Printf("\t- Title: %s\n", rssFeed.Channel.Item[i].Title)
		fmt.Printf("\t- Description: %s\n", rssFeed.Channel.Item[i].Description)
		fmt.Printf("\t- Link: %s\n", rssFeed.Channel.Item[i].Link)
		fmt.Printf("\t- Publish Date: %s\n", rssFeed.Channel.Item[i].PubDate)
	}
}