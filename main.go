package main

import (
	"fmt"

	"github.com/Pranay0205/gator/internal/config"
)




func main() {
		config, err := config.Read()
    if err != nil {
        fmt.Printf("Error reading config: %v\n", err)
        return
    }

		config.SetUser("pranay")

    fmt.Printf("Database URL: %s\n", config)
}