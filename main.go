package main

import (
	"database/sql"
	"log"

	"github.com/TaitA2/bloggregator/internal/config"
	"github.com/TaitA2/bloggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	config, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	db, err := sql.Open("postgres", config.Db_url)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer db.Close()

	dbQueries := database.New(db)

	state := &State{
		config: &config,
		db:     dbQueries}

	commands := GetCommands()

	command, args := GetArgs()

	err = commands.Run(state, Command{command, args})
	if err != nil {
		log.Fatalf("\033[31m[ERROR] %v\n\033[0m", err)
	}

}
