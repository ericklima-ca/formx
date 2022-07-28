package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller interface {
	NewForm(*gin.Context)
}

func NewRouter(c controller) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("static/**/*.html")
	router.Static("/assets", "static/assets/")

	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "New Form",
			})
		})
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		v1.POST("/form_post", c.NewForm)
	}
	return router
}
