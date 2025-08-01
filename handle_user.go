package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Pranay0205/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Login
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	username := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows){
			return fmt.Errorf("user doesn't exists in the database")
		}
		return fmt.Errorf("database error: %v", err)
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("unable to login: %v", err)
	}



	return nil
}


// List Of Users
func handlerUsers(s *state, cmd command) error {

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.Cfg.Username {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}

	return nil
}

// Register User
func handlerRegister(s *state, cmd command) error {
		if len(cmd.Args) != 1 {
			return fmt.Errorf("usage: %v <name>", cmd.Name)
		}

		user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name: cmd.Args[0],
		})

		if err != nil {
			if pqError, ok := err.(*pq.Error); ok{
				if pqError.Code == "23505" {
					return fmt.Errorf("username already exists")
				}
			}
			return fmt.Errorf("failed to create user: %w", err)
		}

		s.Cfg.SetUser(user.Name)

		fmt.Printf("user was created")
		printUser(user)

		return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}