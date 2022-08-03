package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller interface {
	NewForm(*gin.Context)
	ServeStatic(*gin.Context)
	ServeSuccess(*gin.Context)
}

func NewRouter(c controller) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("static/**/*.html")
	router.Static("/assets", "static/assets/")
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/v1/")
	})
	v1 := router.Group("/v1")
	{
		v1.GET("/", c.ServeStatic)
		v1.POST("/form_post", c.NewForm)
		v1.GET("/success", c.ServeSuccess)
	}
	return router
}
