package main

import (
	"log"

	"example.com/medium/db"
	"example.com/medium/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found!")
	}
}

func main() {
	// Setting env variables
	setEnv()

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
