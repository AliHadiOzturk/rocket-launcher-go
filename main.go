package main

import (
	// "context"
	"launcher/config"
	"launcher/manager"
	"launcher/models"
	"launcher/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	socketManager := manager.NewSocketManager()
	go socketManager.Run()

	r := gin.Default()
	r.GET("/ws", socketManager.Handle())

	config.Init()

	rocketManager := manager.NewRocketManager().Init()

	for _, rocket := range rocketManager.Rockets {
		tcpUtil := utils.NewTCPUtil(rocket.Telemetry.Host, rocket.Telemetry.Port)
		go tcpUtil.Connect(func(rocket *models.Rocket) {
			// manager.RedisManager.Client.Publish(context.Background(), "rocket.data", rocket)
			data, err := rocket.Marshal()

			if err != nil {
				return
			}

			socketManager.Publish(data)
		})
	}

	r.Run()
}
