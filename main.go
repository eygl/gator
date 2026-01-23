package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/eygl/gator/internal/config"
	"github.com/eygl/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

  cfg, err := config.Read()
  if err != nil {
		log.Fatalf("error reading config: %v", err)
  }

  db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)

	session_state := state{
		Cfg: &cfg,
		DB: dbQueries,
	}

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
	cmds.register("register", handleRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handleUsersList)
	//cmds.register("command", handleCommand)

	err = cmds.run(&session_state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
