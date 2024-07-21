package main

import (
	"log"
	"net/http"
	"test/layered/internal/handler"
	"test/layered/internal/repository"
	"test/layered/internal/router"
	"test/layered/internal/service"
)

func main() {
	// リポジトリの初期化
	userRepo := repository.NewUserRepository()

	// サービスの初期化
	userService := service.NewUserService(userRepo)

	// ハンドラーの初期化
	userHandler := handler.NewUserHandler(*userService)

	// ルーターの設定
	r := router.SetupRouter(userHandler)

	// サーバーの起動
	log.Println("Starting server on :8082")
	if err := http.ListenAndServe(":8082", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
