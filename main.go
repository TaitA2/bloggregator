package main

import (
	"fmt"
	"os"

	"github.com/TaitA2/bloggregator/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		return
	}
	conf.SetUser("TaitA2")
	homedir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	content, err := os.ReadFile(homedir + "/.gatorconfig.json")
	fmt.Println(string(content))
	fmt.Println(conf)
}
