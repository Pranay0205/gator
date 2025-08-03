package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)


func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	
	client := http.Client{}
	
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed to crate request for rss feed: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := client.Do(req)

	if err != nil {
			return &RSSFeed{}, fmt.Errorf("failed to receive rss feed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return &RSSFeed{}, fmt.Errorf("bad status: %w", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed to read the response body: %w", err)
	}

	var feed RSSFeed
	err = xml.Unmarshal(bodyBytes, &feed)

	if err != nil {
		return &RSSFeed{}, fmt.Errorf("failed to parse xml: %w", err)
	}
		
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Item {
			feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
			feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, nil

}

