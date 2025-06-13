package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	content, err := os.ReadFile(homedir + "/.gatorconfig.json")
	if err != nil {
		return Config{}, err
	}
	var conf Config
	if err := json.Unmarshal(content, &conf); err != nil {
		return Config{}, err
	}
	return conf, nil
}

func (c Config) SetUser(username string) error {
	c.Current_user_name = username
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	js, err := json.Marshal(c)
	os.WriteFile(homedir+"./gatorconfig.json", js, 0666)
	return nil
}
