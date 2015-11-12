package main

import (
	"github.com/gorilla/websocket"
)

// client represent a single chatting user
type client struct {
	socket *websocket.Conn
}
