package main

import (
	"errors"

	"github.com/baletor/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	registered map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registered[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registered[cmd.name]
	if !ok {
		return errors.New("unknown command")
	}
	return f(s, cmd)
}
