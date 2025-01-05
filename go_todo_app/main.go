package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(
		":80",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello %s, World!", r.URL.Path[1:])
		}),
	)
	if err != nil {
		fmt.Printf("failed to start server: %v", err)
	}
}
