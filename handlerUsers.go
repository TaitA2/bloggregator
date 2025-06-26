package main

import (
	"context"
	"fmt"
)

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
