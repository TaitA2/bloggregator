package main

import (
	"fmt"
	"os"

	"github.com/TaitA2/bloggregator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		return
	}

	var state State
	state.config = &config

	commands := GetCommands()

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		fmt.Println(fmt.Errorf("\033[31m[ERROR] Too few arguments!\033[0m"))
		os.Exit(1)
	}

	command, args := cliArgs[1], cliArgs[2:]

	err = commands.Run(&state, Command{command, args})
	if err != nil {
		fmt.Printf("\033[31m[ERROR] %v\n\033[0m", err)
		os.Exit(1)
	}

}
