package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eygl/gator/internal/config"
)

func main() {
  cfg, err := config.Read()
  if err != nil {
		log.Fatalf("error reading config: %v", err)
  }
	fmt.Printf("Read config again: %+v\n", cfg)
	session_state := state{Cfg: &cfg}

	args := os.Args
	if len(os.Args) <= 1 {
		fmt.Printf("Not enough arguments were provided.\n")
		os.Exit(1)
	}
	commandName := args[1]
	commandArgs := args[2:]

	cmd := command {
		Name: commandName,
		Args: commandArgs,
	}
	cmds := commands{Commands: make(map[string]func(*state, command) error)}
	cmds.register("login", handleLogin)
//cmds.register("command", handleCommand)

	err = cmds.run(&session_state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
