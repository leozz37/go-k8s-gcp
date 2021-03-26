package main

import (
	"example.com/medium/db"
	"example.com/medium/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	// Starting DB
	db.ConnectMySQL()

	// Starting Gin router
	r := gin.Default()

	r.GET("/user/:id", handler.FindUser)
	r.GET("/users", handler.FindUsers)
	r.POST("/user", handler.CreateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	r.Run()
}
