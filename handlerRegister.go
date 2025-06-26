package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/TaitA2/bloggregator/internal/database"
	"github.com/google/uuid"
)

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
