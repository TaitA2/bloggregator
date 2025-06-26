package main

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {

	err := s.db.Reset(context.Background())

	if err != nil {
		return fmt.Errorf("Error resetting database: %v", err)
	}

	fmt.Println("Reset database")
	return nil
}
