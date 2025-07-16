package main

import (
	"context"
	"fmt"

	"github.com/TaitA2/bloggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {

		user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
		if err != nil {
			return fmt.Errorf("Error retrieving user %s from database.", s.config.Current_user_name)
		}
		return handler(s, cmd, user)
	}
}
