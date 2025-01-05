package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("Error: %v", err)
	}

	// if err != nil {
	// 	fmt.Printf("failed to start server: %v", err)
	// }
}

func run(ctx context.Context) error {
	s := http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello %s, World!", r.URL.Path[1:])
		}),
	}

	s.ListenAndServe()

	return nil
}
