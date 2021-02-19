package server

import "github.com/gin-gonic/gin"

func StartServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api", func(c *gin.Context) {
		images := [3]string{"/start/trans/img/img11.jpg", "/start/trans/img/img12.jpg", "/start/trans/img/img13.jpg"}
		c.JSON(200, gin.H{
			"images": images,
		})
	})
	r.Static("/start", "photoframe_webclient")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
