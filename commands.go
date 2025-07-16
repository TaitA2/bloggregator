package main

import "fmt"

type Command struct {
	name      string
	arguments []string
}

type Commands struct {
	commandMap map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.commandMap[cmd.name]
	if !ok {
		return fmt.Errorf("Command not found: %s", cmd.name)
	}
	return handler(s, cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.commandMap[name] = f
}

func GetCommands() Commands {
	var commands = Commands{make(map[string]func(*State, Command) error)}

	commands.Register("login", HandlerLogin)
	commands.Register("register", HandlerRegister)
	commands.Register("reset", HandlerReset)
	commands.Register("users", HandlerUsers)
	commands.Register("agg", middlewareLoggedIn(HandlerAgg))
	commands.Register("addfeed", middlewareLoggedIn(HandlerAddFeed))
	commands.Register("feeds", HandlerFeeds)
	commands.Register("follow", middlewareLoggedIn(HandlerFollow))
	commands.Register("unfollow", middlewareLoggedIn(HandlerUnfollow))
	commands.Register("following", middlewareLoggedIn(HandlerFollowing))

	return commands
}
