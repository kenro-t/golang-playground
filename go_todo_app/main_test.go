package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	in := "message"
	resp, err := http.Get("http://localhost:80/" + in)
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
