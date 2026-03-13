package main

import (
	"github.com/31c0cd03/petition-gateway/internal/api/v1/petitions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	petitions.Register(router)

	router.Run("0.0.0.0:8000")
}
