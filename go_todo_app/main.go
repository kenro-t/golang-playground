package main

import (
	"context"
	"fmt"
	"go_todo_app/config"
	"log"
	"net"
	"net/http"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	if err != nil {
		log.Printf("Error: %v", err)
	}
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s, World!", r.URL.Path[1:])
	})

	s := NewServer(l, mux)
	if err := s.Run(context.Background()); err != nil {
		log.Printf("Error: %v", err)
	}
}
