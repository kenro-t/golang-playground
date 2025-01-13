package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func longTaskHandler(w http.ResponseWriter, r *http.Request) {
	// コンテキストを作成
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	// クライアントが切断された場合にキャンセル関数を呼び出す
	notifyChan := make(chan bool)
	go func() {
		<-r.Context().Done()
		cancel()
		notifyChan <- true
	}()

	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			// コンテキストがキャンセルされた場合、処理を中断（ブラウザが閉じられたことも検知可能）
			fmt.Println("Task interrupted")
			return
		default:
			// 長時間処理のシミュレーション
			time.Sleep(1 * time.Second)
			fmt.Println("Processing:", i)
		}
	}
	fmt.Fprintf(w, "Task completed\n")
}

func main() {
	http.HandleFunc("/", longTaskHandler)
	http.ListenAndServe(":80", nil)
}
