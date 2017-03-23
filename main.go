package main

import (
	"log"
	"os"

	"github.com/sasquad92/server-watch/configuration"
	"github.com/sasquad92/server-watch/mail"
	"github.com/sasquad92/server-watch/service"
	"github.com/sasquad92/server-watch/watchdog"
)

var logFile *os.File

func init() {
	// creating log file
	logFile, err := os.OpenFile("watch.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err.Error())
	}

	// setting logger logging destination
	log.SetOutput(logFile)
}
func main() {
	defer logFile.Close()

	config, err := configuration.NewConfigFile("config.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	serv, err := service.NewService(config.ServiceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	notifier, err := mail.NewMail(
		config.NotificationMailTo,
		config.NotificationMailFrom,
		config.MailSmtp,
		config.MailPort,
		config.MailPassword)

	if err != nil {
		log.Fatal(err.Error())
	}

	wd := watchdog.NewWatchdog(
		config.Attemps,
		config.CheckInterval,
		config.RestartInterval,
		serv,
		notifier)

	wd.Watch()
}
