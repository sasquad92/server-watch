package main

import (
	"log"

	"github.com/sasquad92/server-watch/configuration"
	"github.com/sasquad92/server-watch/mail"
	"github.com/sasquad92/server-watch/service"
	"github.com/sasquad92/server-watch/watchdog"
)

func main() {
	config, err := configuration.NewConfigFile("config.json")

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	serv := service.NewService(config.ServiceName)

	notifier := mail.NewMail(
		config.NotificationMailTo,
		config.NotificationMailFrom,
		config.MailSmtp,
		config.MailPort,
		config.MailPassword)

	wd := watchdog.NewWatchdog(
		config.Attemps,
		config.CheckInterval,
		config.RestartInterval,
		serv,
		notifier)

	wd.Watch()
}
