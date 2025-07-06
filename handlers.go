package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/TaitA2/bloggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerAgg(s *State, cmd Command) error {
	user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return err
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for i := range feedFollows {
		url := feedFollows[i].Url

		feed, err := fetchFeed(context.Background(), url)
		if err != nil {
			return fmt.Errorf("Error fetching feed: %v\n", err)
		}

		fmt.Println(feed)
	}
	return nil
}

func HandlerReset(s *State, cmd Command) error {

	err := s.db.Reset(context.Background())

	if err != nil {
		return fmt.Errorf("Error resetting database: %v", err)
	}

	fmt.Println("Reset database")
	return nil
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("No arguments given, expected 1 argument: username")
	}
	username := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err == sql.ErrNoRows || err != nil {
		return fmt.Errorf("User may not exist - %v", err)
	}

	if err := s.config.SetUser(username); err != nil {
		return err
	}
	fmt.Printf("Username set to: %s\n", username)
	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("No arguments given, expected 1 argument: username")
	}

	username := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), username)

	if err != sql.ErrNoRows && err != nil {
		return err
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})

	if err != nil {
		return err
	}

	if err := s.config.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("Created user: %s\n", username)
	fmt.Println(user)

	return nil
}

func HandlerUsers(s *State, cmd Command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error retrieving users from db: %v", err)
	}
	for _, user := range users {
		username := user.Name
		if username == s.config.Current_user_name {
			username += " (current)"
		}
		fmt.Printf("* %s\n", username)
	}
	return nil
}

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("Insuffficient arguments, expected 2: name, url.")
	}
	feedName, feedUrl := cmd.arguments[0], cmd.arguments[1]

	user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return fmt.Errorf("Error retrieving user %s from database.", s.config.Current_user_name)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	})

	HandlerFollow(s, Command{name: "follow", arguments: []string{feedUrl}})
	fmt.Println(feed)

	return nil
}

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error retrieving feeds from database: %v", err)
	}

	fmt.Println(feeds)
	return nil
}

func HandlerFollow(s *State, cmd Command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("No arguments given, expected 1 argument: url")
	}

	url := cmd.arguments[0]
	user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return fmt.Errorf("Error retrieving user %s from database.", s.config.Current_user_name)
	}
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %v\n", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("Error creating feed follow: %v", err)
	}

	fmt.Println(feedFollow)

	return nil

}

func HandlerFollowing(s *State, cmd Command) error {
	user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return fmt.Errorf("Error retrieving user for following command: %v", err)
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.Name_2)
	}
	return nil
}

func HandlerUnfollow(s *State, cmd Command) error {
	return s.db.Unfollow(context.Background(), cmd.arguments[0])
}
