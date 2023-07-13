package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	envPort := "8080"

	if envPort := os.Getenv("PORT"); envPort == "" {
		log.Println("PORT environment variable not set, defaulting	 to 8080")
	}
	log.Println("Starting server on port", envPort)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(http.ListenAndServe(":"+envPort, r))
}
