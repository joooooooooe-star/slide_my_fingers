package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joooooooooe-star/slide_my_fingers/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.POST("/api/songplayed", func(c *gin.Context) {
		handlers.PostSongPlayed(c)
	})

	router.Run()
}
