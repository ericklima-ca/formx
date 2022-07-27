package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ericklima-ca/formx/models"
	"github.com/ericklima-ca/formx/pdf_generator"
)

func NewForm(c *gin.Context) {
	var form models.Form
	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	pdf_generator.BuildPDF(form)
	// TO BE IMPLEMENTED
	c.Redirect(http.StatusFound, "https://www.google.com")
}

// var ctx = context.Background()
//
// func ExampleClient() {
// rdb := redis.NewClient(&redis.Options{
// Addr:     "192.168.15.96:6379",
// Password: "", // no password set
// DB:       0,  // use default DB
// })
//
// err := rdb.Set(ctx, "key", "value", 0).Err()
// if err != nil {
// panic(err)
// }
//
// val, err := rdb.Get(ctx, "key").Result()
// if err != nil {
// panic(err)
// }
// fmt.Println("key", val)
//
// val2, err := rdb.Get(ctx, "key2").Result()
// if err == redis.Nil {
// fmt.Println("key2 does not exist")
// } else if err != nil {
// panic(err)
// } else {
// fmt.Println("key2", val2)
// }
// }
