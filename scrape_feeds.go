package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *State, ctx context.Context) error {
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("Error retrieving next feed to fetch: %v", err)
	}
	feed, err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("Error marking feed as fetched: %v", err)
	}
	rssfeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %v", err)
	}

	for _, item := range rssfeed.Channel.Item {

		fmt.Println(item)
	}
	return nil
}
