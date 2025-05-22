package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// load templates
	r.LoadHTMLGlob("templates/*.tpl")

	r.GET("/page1", func(c *gin.Context) {
		c.HTML(200, "page1.tmpl", gin.H{
			"title": "Page 1 IS HERE",
		})
	})

	r.GET("/page2", func(c *gin.Context) {
		c.HTML(200, "page2.tmpl", gin.H{
			"title": "Page 2 IS HERE",
		})
	})

	r.Run("127.0.0.1:8080")
}
