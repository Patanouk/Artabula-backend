package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/artabula/config", )

	router.GET("/ping", func(c *gin.Context) {
		time.Sleep(1 * time.Second)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run("localhost:8080")
}

func getConfig(c *gin.Context) {

}

type Config struct {

}