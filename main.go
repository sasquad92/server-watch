package main

import (
	"log"

	"github.com/sasquad92/server-watch/configuration"
	"github.com/sasquad92/server-watch/service"
	"github.com/sasquad92/server-watch/watchdog"
)

func main() {
	config, err := configuration.NewConfigFile("config.json")

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	serv := service.NewService(config.ServiceName, config.Path)

	wd := watchdog.NewWatchdog(config.Attemps, config.CheckInterval, config.RestartInterval, serv)

	wd.Watch()
}
