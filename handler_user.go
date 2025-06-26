package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/TaitA2/bloggregator/internal/database"
	"github.com/google/uuid"
)

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
