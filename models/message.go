package models

import "github.com/gorilla/websocket"

type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

type Client struct {
	Conn     *websocket.Conn
	Username string
}

var Clients = make(map[*Client]bool)
