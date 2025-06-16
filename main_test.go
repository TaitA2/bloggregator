package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/TaitA2/bloggregator/internal/config"
)

func TestMain(t *testing.T) {
	conf, err := config.Read()
	if err != nil {
		return
	}
	names := []string{"TestName", "OtherName", "FirstLast"}
	for _, name := range names {
		if !testSetUser(name, conf) {
			t.Errorf("Failed to set username to %s", name)
		}
	}
}

func testSetUser(name string, conf config.Config) bool {
	conf.SetUser(name)
	homedir, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	content, err := os.ReadFile(homedir + "/.gatorconfig.json")
	return string(content) == fmt.Sprintf("{\"db_url\":\"postgres://example\",\"current_user_name\":\"%s\"}", name)
}
