package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ericklima-ca/formx/controllers"
	"github.com/ericklima-ca/formx/pdf_generator"
	"github.com/ericklima-ca/formx/router"
	"github.com/ericklima-ca/formx/services"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var (
		EMAIL_ADDR = os.Getenv("EMAIL_ADDR")
		EMAIL_PASS = os.Getenv("EMAIL_PASS")
		HOST_SMTP  = os.Getenv("HOST_SMTP")
	)

	ms := services.MailerService{
		HostPort: HOST_SMTP,
		User:     EMAIL_ADDR,
		Passcode: EMAIL_PASS,
	}

	controller := controllers.Controller{
		Mailer:       ms,
		PDFGenerator: pdf_generator.PDFGenerator{},
	}
	router := router.NewRouter(controller)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("Listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting.")
}
