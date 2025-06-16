package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/TaitA2/bloggregator/internal/config"
)

func TestSetUser(t *testing.T) {
	conf, err := config.Read()
	if err != nil {
		return
	}
	names := []string{"TestName", "OtherName", "FirstLast"}
	homedir, err := os.UserHomeDir()
	for _, name := range names {
		conf.SetUser(name)
		content, err := os.ReadFile(homedir + "/.gatorconfig.json")
		if err != nil {
			t.Errorf("Failed to read config file while setting user name to: %s", name)
		}
		if string(content) != fmt.Sprintf("{\"db_url\":\"postgres://example\",\"current_user_name\":\"%s\"}", name) {
			t.Errorf("Failed to set username to %s", name)
		}
	}
}
