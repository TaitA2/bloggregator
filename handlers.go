package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/TaitA2/bloggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerHelp(s *State, cmd Command) error {
	fmt.Println(`
Welcome to the Bloggregator Blog Aggregator!

Available commands:

* help                 - prints help message
* register [username]  - register a new user with the given username
* login [username]     - login as the given user
* users                - lists all registered users
* reset                - removes all registered users
* feeds                - lists all available feeds
* addfeed [name] [url] - saves the RSS feed of the given URL as the given name to the available feeds
* follow [feed name]   - adds the named feed to the current user's followed feeds
* unfollow [feed name] - removes the named feed from the current user's followed feeds
* following            - lists all of the current user's followed feeds
* agg [interval]       - aggregates all available feeds once per given interval (1s, 1m, 1h, etc.). Intended to be run in the background.
* browse [limit]       - prints [limit] aggregated posts from followed feeds
			
		`)
	return nil
}

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	var limit int
	if len(cmd.arguments) < 1 {
		limit = 2
	} else {
		arg, err := strconv.Atoi(cmd.arguments[0])
		limit = arg
		if err != nil {
			return fmt.Errorf("Error converting browse limit to integer: %v", err)
		}
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("Error getting posts for user: %v", err)
	}
	for i := range posts {
		fmt.Println(posts[i].Title)
		fmt.Println(posts[i].Description.String)
		fmt.Println()
	}
	return nil
}

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("No arguments given, expected 1 argument: Time between requests (eg. 1s, 1m, 1h, etc.)")
	}
	delay, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Error parsing aggregate duration: %v", err)
	}

	fmt.Printf("Collecting feeds every %v\n", delay)

	ticker := time.NewTicker(delay)

	for ; ; <-ticker.C {
		if err := scrapeFeeds(s, context.Background()); err != nil {
			return fmt.Errorf("Error scraping feeds: %v", err)
		}
	}
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

	user, err := s.db.GetUser(context.Background(), username)

	if err != sql.ErrNoRows && err != nil {
		return err
	}

	user, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
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

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("Insuffficient arguments, expected 2: name, url.")
	}
	feedName, feedUrl := cmd.arguments[0], cmd.arguments[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	})

	if err != nil {
		return err
	}

	HandlerFollow(s, Command{name: "follow", arguments: []string{feedUrl}}, user)
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

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("No arguments given, expected 1 argument: feed name")
	}

	name := cmd.arguments[0]
	feed, err := s.db.GetFeed(context.Background(), name)
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

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.Name_2)
	}
	return nil
}

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	return s.db.Unfollow(context.Background(), cmd.arguments[0])
}
