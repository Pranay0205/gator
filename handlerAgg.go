package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Pranay0205/gator/internal/database"
	"github.com/Pranay0205/gator/internal/rss"
)


func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
    return fmt.Errorf("usage: %s <duration>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])

	if err != nil {
		return fmt.Errorf("invalid duration format. Usage: %s <duration> (examples: 30s, 5m, 1h)", cmd.Name)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <- ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
        fmt.Printf("Error scraping feeds: %v\n", err)
        continue
    }
	}

}


func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get the next feed to fetch: %w", err)
	}

	if nextFeed.LastFetchedAt.Valid {

    fmt.Printf("Fetching feed %q (last fetched %v ago)\n",nextFeed.Name, time.Since(nextFeed.LastFetchedAt.Time))

	} else {
		fmt.Printf("Fetching feed %q (never fetched before)\n", nextFeed.Name)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
				Time:  time.Now().UTC(),
				Valid: true,
			},
		UpdatedAt: time.Now().UTC(),
		ID: nextFeed.ID,
	})

	if err != nil {
		return fmt.Errorf("failed to mark the feed of %v: %w", nextFeed.Name, err)
	}
	
	rssFeed, err := rss.FetchFeed(context.Background(), nextFeed.Url)

	if err != nil {
		return fmt.Errorf("couldn't collect feed %s: %v", nextFeed.Name, err)
	}
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	
	log.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(rssFeed.Channel.Item))

	return nil
}

