package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating new request: %v", err)
	}
	req.Header.Set("User-Agent", "Bloggregator")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error executing http request: %v", err)
	}

	defer res.Body.Close()
	var feed *RSSFeed
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading http response body: %v", err)
	}

	if err := xml.Unmarshal(data, &feed); err != nil {
		return nil, fmt.Errorf("Error unmarshalling xml %v", err)
	}

	decodeEscaped(feed)

	return feed, nil
}

// Function to decode escaped HTML entities (like &ldquo;)
func decodeEscaped(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
}
