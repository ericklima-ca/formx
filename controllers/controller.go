package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ericklima-ca/formx/models"
	pdf "github.com/ericklima-ca/formx/pdf_generator"
)

type mailer interface {
	SendMail([]byte)
}
type pdfGenerator interface {
	BuildPDF(pdf.Data)
}
type Controller struct {
	Mailer       mailer
	PDFGenerator pdfGenerator
}

func (cc Controller) NewForm(c *gin.Context) {
	var form models.Form

	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	cc.PDFGenerator.BuildPDF(form)

	jsonBytes, errJson := json.Marshal(gin.H{
		"To":       form.Email,
		"Subject":  "Form from " + form.Name,
		"Body":     "Data submitted by " + form.Name,
		"Customer": form.Name,
	})
	if errJson != nil {
		log.Fatal(errJson)
	}
	cc.Mailer.SendMail(jsonBytes)

	// TO BE IMPLEMENTED
	c.Redirect(http.StatusFound, "https://www.google.com")
}
