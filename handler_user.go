package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username required")
	}
	username := cmd.args[0]
	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("User %s has been set\n", username)
	return nil
}
