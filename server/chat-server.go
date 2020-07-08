package server

import (
	"chatroom/core"
	"github.com/gin-gonic/gin"
)

var Room = core.NewRoom()

func NewServer() *gin.Engine {
	s := gin.Default()
	// static files
	s.Static("/static", "./static")
	s.StaticFile("/", "web/index.html")
	s.StaticFile("/chatroom", "./web/chatroom.html")

	// websocket
	s.GET("/ws/socket", Websocket.Handle())

	return s
}
