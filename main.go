package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hlts2/gin-server-template/interfaces"
)

func main() {
	g := gin.Default()

	g.GET("/tweets/", interfaces.GetTweets)
	g.GET("/tweets/:id", interfaces.GetTweet)

	var port string
	if port = os.Getenv("PORT"); 0 == len(port) {
		port = "8080"
	}

	if err := g.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
