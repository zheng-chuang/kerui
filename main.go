package main

import (
	"net/http"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.HTMLRender = gintemplate.Default()
	r.Static("assets", "./assets")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title": "Gin框架",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
