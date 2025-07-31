package main

import (
	"log"
	"os"

	"github.com/Pranay0205/gator/internal/config"
)




func main() {
		config, err := config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }

    appState := state{Cfg: config}

    cmdHandler := commands{commandHandler: make(map[string]func(*state, command) error)}

    cmdHandler.register("login", handlerLogin)

    args := os.Args

    if len(args) < 2 {
      log.Fatal("Usage: cli <command> [args...]")
    } 

    cmdName, arguments := args[1], args[2:] 

    cmd := command{Name: cmdName, Args: arguments}

    err = cmdHandler.run(&appState, cmd)
    
    if err != nil {
      log.Fatal(err)
    }


}