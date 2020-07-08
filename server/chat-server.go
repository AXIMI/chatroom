package server

import (
	"chatroom/core"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var Room = core.NewRoom()

type User struct {
	Id       int
	Username string
	Password string
}


func NewServer() *gin.Engine {
	s := gin.Default()


	// static files
	s.Static("/static", "./static")
	s.StaticFile("/", "web/index.html")
	s.StaticFile("/chatroom", "./web/chatroom.html")

	s.POST("/authentication", func(c *gin.Context){
		db, err := gorm.Open("mysql", "root:12345678@/chat?charset=utf8&parseTime=True&loc=Local")
		defer db.Close()

		if err != nil {
			panic(err)
		}

		decoder := json.NewDecoder(c.Request.Body)
		// 解析参数 存入map
		var params map[string]string
		decoder.Decode(&params)
		username := params["username"]
		password := params["password"]

		// 检查数据库中是否有该用户
		var user User
		db.Where("username = ?", username).Find(&user)
		if user.Username == "" {
			c.JSON(200, gin.H{
				"status":  1,
				"message": "there is no such user",
			})
		}else if user.Password != password {
			c.JSON(200, gin.H{
				"status":  2,
				"message": "wrong password",
			})
		}else {
			c.JSON(200, gin.H{
				"status":  0,
				"message": "access",
			})
		}
	})

	s.POST("/sign_up", func(c *gin.Context){
		db, err := gorm.Open("mysql", "root:12345678@/chat?charset=utf8&parseTime=True&loc=Local")
		defer db.Close()

		if err != nil {
			panic(err)
		}

		decoder := json.NewDecoder(c.Request.Body)
		// 解析参数 存入map
		var params map[string]string
		decoder.Decode(&params)
		username := params["username"]
		password := params["password"]

		// 检查数据库中是否有该用户
		var user User
		db.Where("username = ?", username).Find(&user)
		if user.Username == "" {
			// 将该数据存储到数据库
			db.Create(&User{Username: username, Password: password})

			c.JSON(200, gin.H{
				"status":  0,
				"message": "sign up successfully",
			})
		} else {
			c.JSON(200, gin.H{
				"status":  1,
				"message": "this account already exists",
			})
		}
	})

	// websocket
	s.GET("/ws/socket", Websocket.Handle())

	return s
}
