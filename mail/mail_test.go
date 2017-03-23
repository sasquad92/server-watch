package mail_test

import (
	"testing"

	"github.com/sasquad92/server-watch/mail"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestNewMail_IncorrectToAddress(t *testing.T) {
	to := "wrondAddress"
	from := "ex@pmle.com"
	smtp := "ex.am.com"
	port := 123
	pass := "secret"

	t.Log("Recipient email address should have '@'")

	mail, err := mail.NewMail(to, from, smtp, port, pass)

	if mail != nil {
		t.Error("Mail instance should be nil", ballotX)
	} else {
		t.Log("Mail instance should be nil", checkMark)
	}

	if err == nil {
		t.Error("Error object should not be nil", ballotX)
	} else {
		t.Log("Error object should not be nil", checkMark)
	}
}

func TestNewMail_IncorrectFromAddress(t *testing.T) {
	to := "ex@mple.com"
	from := "wrondAddress"
	smtp := "ex.am.com"
	port := 123
	pass := "secret"

	t.Log("Adressee email address should have '@'")

	mail, err := mail.NewMail(to, from, smtp, port, pass)

	if mail != nil {
		t.Error("Mail instance should be nil", ballotX)
	} else {
		t.Log("Mail instance should be nil", checkMark)
	}

	if err == nil {
		t.Error("Error object should not be nil", ballotX)
	} else {
		t.Log("Error object should not be nil", checkMark)
	}
}

func TestNewMail_CorrectOutput(t *testing.T) {
	to := "ex@mple.cpm"
	from := "ex@pmle.com"
	smtp := "ex.am.pl"
	port := 123
	pass := "secret"

	mail, err := mail.NewMail(to, from, smtp, port, pass)

	if mail != nil {
		t.Log("Mail instance should not be nil", checkMark)
	} else {
		t.Error("Mail instance should not be nil", ballotX)
	}

	if err == nil {
		t.Log("Error object should be nil", checkMark)
	} else {
		t.Error("Error object should be nil", ballotX)
	}
}
