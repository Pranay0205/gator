package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Pranay0205/gator/internal/config"
	"github.com/Pranay0205/gator/internal/database"

	_ "github.com/lib/pq"
)




func main() {
		config, err := config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }

    db, err := sql.Open("postgres", config.DbURL)
    if err != nil{
      log.Fatalf("error connecting to the database: %v", err)
    }

    dbQueries := database.New(db)


    appState := state{Cfg: config, db: dbQueries}

    cmdHandler := commands{commandHandler: make(map[string]func(*state, command) error)}

    cmdHandler.register("login", handlerLogin)

    cmdHandler.register("register", handlerRegister)
    
    cmdHandler.register("reset", handlerReset)

    cmdHandler.register("users", handlerUsers)

    cmdHandler.register("user", handlerUser)

    cmdHandler.register("agg", handlerAgg)

    cmdHandler.register("addfeed", handlerAddFeed)

    cmdHandler.register("feeds", handlerListFeeds)

    cmdHandler.register("follow", handlerFollowFeed)

    cmdHandler.register("following", handlerFollowFeedForUser)
    
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