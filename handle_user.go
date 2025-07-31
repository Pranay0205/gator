package main

import "fmt"


func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no arguements provided")
	}

	username := cmd.Args[0]

	err := s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("unable to login: %v", err)
	}

	fmt.Printf("user %s has been set\n", username)

	return nil
}
