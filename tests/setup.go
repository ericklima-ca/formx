package tests

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/ericklima-ca/formx/controllers"
	"github.com/ericklima-ca/formx/pdf_generator"
	"github.com/ericklima-ca/formx/router"
	"github.com/gin-gonic/gin"
)

func Init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func setupRouter() *gin.Engine {
	controller := controllers.Controller{
		Mailer:       MailerMock{},
		PDFGenerator: pdf_generator.PDFGenerator{},
	}
	r := router.NewRouter(controller)
	return r
}

type MailerMock struct{}

func (m MailerMock) SendMail(_b []byte) {
	log.Println("Email sent successfully.")
}
