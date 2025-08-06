package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Pranay0205/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
			return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: cmd.Args[0],
		Url: cmd.Args[1],
		UserID: user.ID,
	})

	if err != nil {
		return fmt.Errorf("failed to add feed to the database: %w", err)
	}

	followedFeed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	
	if err != nil {
		return fmt.Errorf("failed to follow the feed: %w", err)
	}

	fmt.Println("Feed was added to db successfully!")
	fmt.Printf("%v is now following %v\n", followedFeed.UserName, followedFeed.FeedName)


	printFeed(feed)
	
	return nil
}


func handlerListFeeds(s *state, cmd command) error {
	feedRows, err := s.db.GetFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("coudn't get the feed from database: %w", err)
	}
	separator := strings.Repeat("=", 50)
	
	for _, row := range feedRows {
			fmt.Printf(
				"Feed Name: %s\nFeed URL: %s\nUsername: %s\n%s\n",
				row.FeedName,
				row.Url,
				row.UserName,
				separator,
			)
		}

	return nil
	
}


func printFeed(feed database.Feed) {
	fmt.Printf(" * %-15s %v\n", "ID:", feed.ID)
	fmt.Printf(" * %-15s %v\n", "Name:", feed.Name)
	fmt.Printf(" * %-15s %v\n", "URL:", feed.Url)
	fmt.Printf(" * %-15s %v\n", "Created Date:", feed.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf(" * %-15s %v\n", "Updated Date:", feed.UpdatedAt.Format("2006-01-02 15:04:05"))
}

