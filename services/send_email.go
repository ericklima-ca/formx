package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"

	"gopkg.in/mail.v2"
)

type MailerService struct {
	HostPort string
	User     string
	Passcode string
}

type message struct {
	To       string `json:"to,omitempty"`
	Subject  string `json:"subject,omitempty"`
	Body     string `json:"body,omitempty"`
	Customer string `json:"customer,omitempty"`
}

func (ms MailerService) SendMail(jsonBody []byte) {
	var msg message

	if err := json.Unmarshal(jsonBody, &msg); err != nil {
		log.Fatal(err)
	}

	m := mail.NewMessage()

	m.SetHeader("From", ms.User)
	m.SetHeader("To", msg.To)
	m.SetHeader("Subject", msg.Subject)
	filename := fmt.Sprintf("%s.pdf", msg.Customer)
	m.Attach("temp/"+filename, mail.Rename(filename))

	m.SetBody("text/html", msg.Body)
	host, port_str, _ := net.SplitHostPort(ms.HostPort)
	port_number, _ := strconv.Atoi(port_str)
	d := mail.NewDialer(host, port_number, ms.User, ms.Passcode)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}
