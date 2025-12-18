package main

import "errors"

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

