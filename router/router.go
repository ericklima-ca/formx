package router

import (
	"net/http"

	"github.com/ericklima-ca/formx/controllers"
	"github.com/gin-gonic/gin"
)

var data = []gin.H{
	{"Name": "Erick", "Value": 1},
	{"Name": "Amorim", "Value": 2},
}

type Router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("static/**/*.html")
	router.Static("/assets", "static/assets/")

	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"PageTitle": "Hello World",
				"Data":      data,
			})
		})
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		v1.POST("/form_post", controllers.NewForm)
	}
	return router
}
