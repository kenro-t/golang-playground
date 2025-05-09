package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Handler: mux,
		},
		l: l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// グレースフルシャットダウンのためのコンテキスト
	// 割り込み（Interrupt）があったときに終了（SIGTERM）を実行する
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	// errgroupでサーバーを起動する（errgroupは通常のゴルーチンと違いエラーを返してくれる）
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close %v", err)
			return err
		}
		return nil
	})

	// シグナルを受け取ったらシャットダウンする
	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown %v", err)
		return err
	}
	// グレースフルシャットダウン待ち
	return eg.Wait()
}
