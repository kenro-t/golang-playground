package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestServer_Run(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s, World!", r.URL.Path[1:])
	})
	eg.Go(func() error {
		s := NewServer(l, mux)
		return s.Run(ctx)
	})

	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	t.Logf("try request to: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("failed to get %v", err)
	}
	defer resp.Body.Close()

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed read body %v", err)
	}

	want := fmt.Sprintf("Hello %s, World!", in)
	if string(got) != want {
		t.Errorf("want %s, got %s", want, got)
	}

	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
