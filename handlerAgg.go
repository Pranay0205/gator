package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Pranay0205/gator/internal/database"
	"github.com/Pranay0205/gator/internal/rss"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
	for ; ; <-ticker.C {
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

		fmt.Printf("Fetching feed %q (last fetched %v ago)\n", nextFeed.Name, time.Since(nextFeed.LastFetchedAt.Time))

	} else {
		fmt.Printf("Fetching feed %q (never fetched before)\n", nextFeed.Name)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		UpdatedAt: time.Now().UTC(),
		ID:        nextFeed.ID,
	})

	if err != nil {
		return fmt.Errorf("failed to mark the feed of %v: %w", nextFeed.Name, err)
	}

	rssFeed, err := rss.FetchFeed(context.Background(), nextFeed.Url)

	if err != nil {
		return fmt.Errorf("couldn't collect feed %s: %v", nextFeed.Name, err)
	}

	err = saveRSSItem(s, *rssFeed, nextFeed.ID)

	if err != nil {
		return fmt.Errorf("failed to save feed %v", err)
	}

	log.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(rssFeed.Channel.Item))

	return nil
}

func saveRSSItem(s *state, rssFeed rss.RSSFeed, feedID uuid.UUID) error {
	for _, item := range rssFeed.Channel.Item {
		if item.Link == "" {
			log.Printf("Skipping post with empty URL: %s", item.Title)
			continue
		}

		createdPost, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: parsePubDate(item.PubDate),
			FeedID:      feedID,
		})

		if err != nil {
			if pqError, ok := err.(*pq.Error); ok {
				if pqError.Code == "23505" {
					fmt.Printf("given feed already exists %s", item.Title)
					continue
				}
			}
			log.Printf("Error creating post: %v", err)
			continue
		}
		fmt.Printf("Saved post: %s, published: %s\n",
			createdPost.Title,
			createdPost.PublishedAt.Format("Jan 2, 2006 at 3:04 PM"))
	}

	return nil
}

func parsePubDate(dateStr string) time.Time {
	formats := []string{
		"Mon, 02 Jan 2006 15:04:05 MST",   // RFC822 - most common
		"Mon, 02 Jan 2006 15:04:05 -0700", // RFC822 with numeric timezone
		"2006-01-02T15:04:05Z",            // RFC3339
		"2006-01-02T15:04:05-07:00",       // RFC3339 with timezone
		"2006-01-02 15:04:05",             // Your current format
		"02 Jan 2006 15:04:05 MST",        // Without weekday
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t.UTC()
		}
	}

	return time.Now().UTC()
}
