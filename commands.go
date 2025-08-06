package main

import (
	"fmt"

	"github.com/Pranay0205/gator/internal/config"
	"github.com/Pranay0205/gator/internal/database"
)

type state struct {
	db *database.Queries
	Cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

// TO DO - FOR SHOWING ALL COMMANDS WITH DESCRIPTION REFACTORING OF CODE IS NEEDED
// type commandDetails struct {
// 	handler func(*state, command) error
// 	description string	
// }

type commands struct {
	commandHandler map[string]func(*state, command) error
}


func (c *commands) run(s *state, cmd command) error {
	cmdfunc, exists := c.commandHandler[cmd.Name]

	if !exists {
		return fmt.Errorf("command provided doesn't exists")
	} 

	err := cmdfunc(s, cmd)
	if err != nil {
		return fmt.Errorf("failed to execute command %q: %w", cmd.Name, err)
	}

	return nil
}

func (c *commands) getHelp() error {
		if len(c.commandHandler) == 0 {
			return fmt.Errorf("no commands registered")
		}
		fmt.Println("\nPlease use below commands for using gator:")
		for commandName := range c.commandHandler {
			fmt.Printf("* Command: %s\n", commandName)
		}

		return nil
}


func (c *commands) register(name string, f func(*state, command) error) {
    if len(name) == 0 {
        panic("programming error: command name cannot be empty during registration")
    }

		c.commandHandler[name] = f
}
