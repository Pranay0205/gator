package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)


func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no arguements provided")
	}

	username := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), username)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows){
			log.Fatalf("user doesn't exists in the database")
		}
		return fmt.Errorf("database error: %v", err)
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("unable to login: %v", err)
	}

	fmt.Printf("user %s has been set\n", username)

	return nil
}
