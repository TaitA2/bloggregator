package main

import "fmt"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("No arguments given, expected 1 argument: username")
	}
	username := cmd.arguments[0]
	if err := s.config.SetUser(username); err != nil {
		return err
	}
	fmt.Printf("Username set to: %s\n", username)
	return nil
}
