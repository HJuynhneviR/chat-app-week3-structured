package main

import (
	"chat/controllers"
	"chat/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitRedis()

	r := gin.Default()
	r.GET("/ws/:username", controllers.HandleWebSocket)
	r.GET("/online", controllers.ListOnlineUsers)
	r.GET("/history/:username", controllers.GetHistory)

	go controllers.HandleMessages()

	r.Run(":8080")
}
