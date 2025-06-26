package main

import (
	"fmt"
	"os"
)

func GetArgs() (command string, args []string) {
	cliArgs := os.Args

	if len(cliArgs) < 2 {
		fmt.Println(fmt.Errorf("\033[31m[ERROR] Too few arguments!\033[0m"))
		os.Exit(1)
	}

	command, args = cliArgs[1], cliArgs[2:]
	return
}
