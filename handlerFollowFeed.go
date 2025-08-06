package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Pranay0205/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	feed, err := s.db.GetFeed(context.Background(), strings.TrimSpace(cmd.Args[0])) 

	if err != nil {
		return fmt.Errorf("coudn't get the feed details: %w", err)
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

	fmt.Printf("%v is now following %v\n", followedFeed.UserName, followedFeed.FeedName)

	return nil

}



func handlerFollowFeedForUser(s *state, cmd command, user database.User) error {

	followFeedRows, err := s.db.GetFollowFeedForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("coudn't get the follow feed of the user %v: %w", user.Name, err)
	}

	fmt.Println("User follows below feeds:")
	for i, feed := range followFeedRows {
		fmt.Printf("%d) %v\n", i + 1, feed.Name)
	}

	return nil
}

