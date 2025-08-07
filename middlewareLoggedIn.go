package main

import (
	"context"
	"fmt"

	"github.com/Pranay0205/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {

		user, err := s.db.GetUser(context.Background(), s.Cfg.Username)

		if err != nil {
			return fmt.Errorf("couldn't get the current user details: %w", err)
		}

		err = handler(s, cmd, user)
		if err != nil {
			return fmt.Errorf("coudn't complete the operation: %w", err)
		}

		return nil
	}

}
