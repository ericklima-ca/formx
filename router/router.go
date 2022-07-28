package router

import (
	"github.com/gin-gonic/gin"
)

type controller interface {
	NewForm(*gin.Context)
	ServeStatic(*gin.Context)
}

func NewRouter(c controller) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("static/**/*.html")
	router.Static("/assets", "static/assets/")

	v1 := router.Group("/v1")
	{
		v1.GET("/", c.ServeStatic)
		v1.POST("/form_post", c.NewForm)
	}
	return router
}
