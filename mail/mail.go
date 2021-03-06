package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Mail struct {
	mailTo   string
	mailFrom string
	smtp     string
	port     int
	pass     string
}

// NewMail creates new instance of Mail and returns a reference to it
func NewMail(to string, from string, smtp string, port int, pass string) (*Mail, error) {

	if !strings.Contains(to, "@") {
		return nil, fmt.Errorf("Incorrect recipment email address.")
	}
	if !strings.Contains(from, "@") {
		return nil, fmt.Errorf("Incorrect addressee email address.")
	}
	if !(len(pass) > 0) {
		return nil, fmt.Errorf("Email password is incorrect.")
	}

	mail := Mail{
		mailTo:   to,
		mailFrom: from,
		smtp:     smtp,
		port:     port,
		pass:     pass,
	}
	return &mail, nil
}

// Send sends an email at specified address
func (m *Mail) Send(subject string, body string) error {

	msg := "From: " + m.mailFrom + "\n" + "To: " + m.mailTo + "\n" +
		"Subject: " + subject + "\n\n" + body

	port := fmt.Sprintf(":%d", m.port)

	err := smtp.SendMail(m.smtp+port, smtp.PlainAuth("", m.mailFrom, m.pass, m.smtp), m.mailFrom, []string{m.mailTo}, []byte(msg))

	if err != nil {
		return err
	}

	return nil
}
