package main

import (
	"fmt"

	"github.com/eygl/gator/internal/config"
)

type commands struct {
	Commands map[string]func(*state, command) error
}

type command struct {
	Name		string
	Args		[]string
}

type state struct {
	Cfg *config.Config
}

func handleLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("No username given.")
	}
	username := cmd.Args[0]
	err := s.Cfg.SetUser(username)
	if err != nil {
		return err	
	}
	fmt.Printf("User %s has logged in.\n", username)
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	err := c.Commands[cmd.Name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Commands[name] = f
}
