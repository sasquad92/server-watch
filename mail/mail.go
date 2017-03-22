package mail

import (
	"fmt"
	"log"
	"net/smtp"
)

type Mail struct {
	mailTo   string
	mailFrom string
	smtp     string
	port     int
	pass     string
}

// NewMail creates new instance of Mail and returns a reference to it
func NewMail(to string, from string, smtp string, port int, pass string) *Mail {
	mail := Mail{
		mailTo:   to,
		mailFrom: from,
		smtp:     smtp,
		port:     port,
		pass:     pass,
	}
	return &mail
}

// Send sends an email at specified address
func (m *Mail) Send(subject string, body string) error {

	msg := "From: " + m.mailFrom + "\n" + "To: " + m.mailTo + "\n" +
		"Subject: " + subject + "\n\n" + body

	port := fmt.Sprintf(":%d", m.port)

	err := smtp.SendMail(m.smtp+port, smtp.PlainAuth("", m.mailFrom, m.pass, m.smtp), m.mailFrom, []string{m.mailTo}, []byte(msg))

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
