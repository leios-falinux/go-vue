package main

import "net/http"
import "github.com/gin-gonic/gin"
import "github.com/gobuffalo/packr"

func main() {
	box := packr.NewBox("html")

	r := gin.Default()
	r.StaticFS("/", http.FileSystem(box))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
