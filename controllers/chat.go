package controllers

import (
	"chat/middlewares"
	"chat/models"
	"chat/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var broadcast = make(chan models.Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context) {
	username := c.Param("username")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	client := &models.Client{Conn: conn, Username: username}
	models.Clients[client] = true
	utils.Rdb.SAdd(context.Background(), "online", username)

	defer func() {
		delete(models.Clients, client)
		utils.Rdb.SRem(context.Background(), "online", username)
		conn.Close()
	}()

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		if !middlewares.Limiter.Allow() {
			conn.WriteJSON(gin.H{"error": "Rate limit exceeded"})
			continue
		}

		msg.Sender = username
		key := fmt.Sprintf("chat:%s", strings.ToLower(username))
		utils.Rdb.RPush(context.Background(), key, msg.Content)
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range models.Clients {
			client.Conn.WriteJSON(msg)
		}
	}
}

func ListOnlineUsers(c *gin.Context) {
	users, _ := utils.Rdb.SMembers(context.Background(), "online").Result()
	c.JSON(200, gin.H{"online": users})
}

func GetHistory(c *gin.Context) {
	username := c.Param("username")
	key := fmt.Sprintf("chat:%s", strings.ToLower(username))
	history, _ := utils.Rdb.LRange(context.Background(), key, 0, -1).Result()
	c.JSON(200, gin.H{"history": history})
}
