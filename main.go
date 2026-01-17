package main

import (
	"fmt"
	"log"

	"github.com/eygl/gator/internal/config"
)

func main() {
  cfg, err := config.Read()
  if err != nil {
		log.Fatalf("error reading config: %v", err)
  }
	fmt.Printf("Read config again: %+v\n", cfg)

	err = cfg.SetUser("erick")
  if err != nil {
		log.Fatalf("error setting user in config: %v", err)
  }
	fmt.Printf("Set user to %s: %+v\n", "erick", cfg)

  cfg, err = config.Read()
  if err != nil {
		log.Fatalf("error reading config: %v", err)
  }
	fmt.Printf("Read config again: %+v\n", cfg)
}
