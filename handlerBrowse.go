package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {
	limit := 2

	if len(cmd.Args) > 0 {

		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: must be a number")
		}
		if parsedLimit <= 0 {
			return fmt.Errorf("limit must be greater than 0")
		}
		limit = parsedLimit
	}

	limit32 := int32(limit)

	posts, err := s.db.GetPosts(context.Background(), limit32)
	if err != nil {
		return fmt.Errorf("couldn't fetch posts: %w", err)
	}
	for i, post := range posts {
		fmt.Printf("=== Post %d ===\n", i+1)
		fmt.Printf("Feed Name: %s\n", post.FeedName)
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Published: %s\n", post.PublishedAt.Format("Jan 2, 2006"))
		fmt.Printf("URL: %s\n", post.Url)

		// Handle nullable description
		if post.Description.Valid && post.Description.String != "" {
			fmt.Printf("Description: %s\n", post.Description.String)
		}
		fmt.Printf("\n") // spacing between posts
	}

	return nil
}
