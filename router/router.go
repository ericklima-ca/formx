package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var data = []gin.H{
	{"Name": "Erick", "Value": 1},
	{"Name": "Amorim", "Value": 2},
}

type Router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func NewRouter(pathToStatic string) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob(pathToStatic + "/**/*.html")
	router.Static("/assets", pathToStatic+"/assets/")

	v1 := router.Group("/api/v1")
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
	}
	return router
}
