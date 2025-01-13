package main

import (
	"context"
	"fmt"
	"go_todo_app/config"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("Error: %v", err)
	}
}

func run(ctx context.Context) error {
	// グレースフルシャットダウンのためのコンテキスト
	// 割り込み（Interrupt）があったときに終了（SIGTERM）を実行する
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 設定ファイルを読み込む
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// ポートをリッスンする
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d, %v", cfg.Port, err)
	}

	s := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(5 * time.Second)
			fmt.Fprintf(w, "Hello %s, World!", r.URL.Path[1:])
		}),
	}

	// errgroupでサーバーを起動する（errgroupは通常のゴルーチンと違いエラーを返してくれる）
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close %v", err)
			return err
		}
		return nil
	})

	// シグナルを受け取ったらシャットダウンする
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown %v", err)
		return err
	}

	return eg.Wait()
}
