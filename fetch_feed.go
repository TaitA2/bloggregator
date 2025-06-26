package main

import (
	"context"
	"encoding/xml"
	"fmt"
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

	return feed, nil
}
