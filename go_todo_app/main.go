package main

import (
	"context"
	"go_todo_app/config"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	cfg.Port
	s := NewServer(l, mux)
	if err := Run(context.Background()); err != nil {
		log.Printf("Error: %v", err)
	}
}
