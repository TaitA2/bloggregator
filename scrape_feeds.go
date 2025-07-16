package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/TaitA2/bloggregator/internal/database"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
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
		desc := sql.NullString{String: item.Description, Valid: true}
		if len(item.Description) < 1 {
			desc = sql.NullString{String: item.Description, Valid: false}
		}

		pubDate, err := dateparse.ParseAny(item.PubDate)
		if err != nil {
			return fmt.Errorf("Error parsing publish date: %v", err)
		}

		_, err = s.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         feed.Url,
			Description: desc,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})

		if err != nil && !strings.Contains(err.Error(), `duplicate key value violates unique constraint "posts_url_key"`) {
			return fmt.Errorf("Error saving post to database: %v", err)
		}
	}
	return nil
}
