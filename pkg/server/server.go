package server

import (
	"github.com/gin-gonic/gin"
	"github.com/schlunsen/gopexels/pkg/pexels"
	"github.com/spf13/viper"
)

func StartServer(query string) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api", func(c *gin.Context) {
		client := pexels.NewClient(viper.GetString("APIKEY"))

		images := client.SimpleQuery(query, 5)
		c.JSON(200, gin.H{
			"images": images,
		})
	})
	r.Static("/start", "photoframe_webclient")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
