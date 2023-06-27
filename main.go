package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"request-handler-unit-test-example/users"
)

func main() {
	router := gin.Default()
	requestHandler := users.RequestHandler{}

	router.GET("/users", requestHandler.GetUsers)
	router.POST("/users", requestHandler.CreateUser)

	err := router.Run(":8080")
	if err != nil {
		log.Println(err)
	}
}
