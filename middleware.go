package main

import (
	"context"
	"log"

	"github.com/TaitA2/bloggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	user, err := handler.s.db.GetUser(context.Background(), handler.s.config.Current_user_name)
	if err != nil {
		return nil

		log.Fatalf("Error retrieving user %s from database.", handler.s.config.Current_user_name)
	}
	return nil
}
