package main

import "net/http"
import "github.com/gin-gonic/gin"
import "github.com/gobuffalo/packr"
import "path"

func BoxServe(urlPrefix string, box packr.Box) gin.HandlerFunc {
	filesystem := http.FileSystem(box)
	fileserver := http.FileServer(filesystem)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		if box.Has(path.Join(urlPrefix, c.Request.URL.Path)) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

func main() {
	box := packr.NewBox("html")

	r := gin.Default()

	r.Use(BoxServe("/", box))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
