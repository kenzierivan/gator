package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/kenzierivan/gator/internal/database"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	commands map[string]func(*state, command)error
}

func (c *commands) run(s *state, cmd command) error {
	command, ok := c.commands[cmd.Name]
	if !ok {
		return errors.New("command does not exist")
	}
	return command(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func middlewareLoggedIn(handler func(s* state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
		if err != nil {
			return fmt.Errorf("couldn't retrieve user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
