package main

import (
	"log"
	"os"
)

func GetArgs() (string, []string) {
	cliArgs := os.Args

	if len(cliArgs) < 2 {
		log.Fatalf("\033[31m[ERROR] Too few arguments!\033[0m")
	}

	return cliArgs[1], cliArgs[2:]
}
