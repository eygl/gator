package main

import (
	"context"
	"fmt"
	"time"

	"github.com/eygl/gator/internal/config"
	"github.com/eygl/gator/internal/database"
	"github.com/google/uuid"
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
	DB  *database.Queries
}

func handleLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("No username given.")
	}
	username := cmd.Args[0]
	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("User %s does not exists.", username)
	}

	err = s.Cfg.SetUser(username)
	if err != nil {
		return err	
	}
	return nil
}

func handleRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("No username given.")
	}

	username := cmd.Args[0]
	user, err := s.DB.GetUser(context.Background(), username)
	if user.Name == username {
		return fmt.Errorf("User %s is already registered.\n", username)
	}

	userParams := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),	
		UpdatedAt: time.Now(),	
		Name: username,
	}

	_, err = s.DB.CreateUser(context.Background(),userParams)
	err = s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Query error. Could not register user.")
	}
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
