package main

import (
	"fmt"

	"github.com/sasquad92/server-watch/configuration"
	"github.com/sasquad92/server-watch/watchdog"
)

func main() {
	config, err := configuration.NewConfigFile("config.json")

	if err != nil {
		fmt.Println("ERROR!", err)
		return
	}

	watchdog := watchdog.NewWatchdog(config.Attemps, config.CheckInterval, config.RestartInterval)

	attemps := watchdog.GetAttemps()

	fmt.Println("watchdog will try to restart", config.ServiceName, "up to", attemps, "times")
}
