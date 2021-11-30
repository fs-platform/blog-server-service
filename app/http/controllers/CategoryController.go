package controllers

import "github.com/gin-gonic/gin"

type Category struct {
	Controller
}

func (c *Category) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
