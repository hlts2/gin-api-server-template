package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hlts2/gin-server-template/interfaces"
)

func main() {
	g := gin.Default()

	g.GET("/Users/", interfaces.GetUsers)
	g.GET("/Users/:id", interfaces.GetUser)

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "8080"
	}

	if err := g.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
