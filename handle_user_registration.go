package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Pranay0205/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handleRegister(s *state, cmd command) error {
		if len(cmd.Args) == 0 || len(cmd.Args[0]) == 0 {
			return fmt.Errorf("no user name provided")
		}

		user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(), Name: cmd.Args[0]})

		if err != nil {
			if pqError, ok := err.(*pq.Error); ok{
				if pqError.Code == "23505" {
					log.Fatalf("username already exists")
				}
			}
			return fmt.Errorf("failed to create user: %v", err)
		}

		s.Cfg.SetUser(user.Name)

		fmt.Printf("user was created")
		log.Printf("Created user - ID: %s, Name: %s, CreatedAt: %s", user.ID, user.Name, user.CreatedAt)

		return nil
}