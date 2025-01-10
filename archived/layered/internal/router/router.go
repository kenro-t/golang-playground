package router

import (
	"test/layered/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/users/:id", userHandler.GetUser)
	return router
}
