package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/TaitA2/bloggregator/internal/config"
	"github.com/TaitA2/bloggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	config, err := config.Read()
	if err != nil {
		return
	}

	db, err := sql.Open("postgres", config.Db_url)

	dbQueries := database.New(db)

	var state State
	state.config = &config
	state.db = dbQueries

	commands := GetCommands()

	command, args := GetArgs()

	err = commands.Run(&state, Command{command, args})
	if err != nil {
		fmt.Printf("\033[31m[ERROR] %v\n\033[0m", err)
		os.Exit(1)
	}

}
