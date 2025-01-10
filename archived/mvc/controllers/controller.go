package controllers

import "github.com/gin-gonic/gin"

type Controller struct {
}

func (c *Controller) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Yahoao!!!",
	})
}
