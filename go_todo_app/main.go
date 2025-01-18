package main

import (
	"context"
	"fmt"
	"go_todo_app/config"
	"log"
	"net"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Printf("Error: %v", err)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("Server start with %s", url)

	mux := NewMux()
	s := NewServer(l, mux)

	return s.Run(ctx)
}
