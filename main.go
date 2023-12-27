package main

import (
	"github.com/gin-gonic/gin"
	"launcher/helpers"
	"launcher/manager"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	eventHelper := helpers.NewEventHelper()

	manager.NewRocketMAnager().Init().ConnectRockets(eventHelper)

	r.Run()
}
