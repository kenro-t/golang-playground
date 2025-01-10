package routes

import (
	"mvc/controllers"

	"github.com/gin-gonic/gin"
)

var c *controllers.Controller = &controllers.Controller{}

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		c.Index(ctx)
	})
	return r
}
