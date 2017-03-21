package watchdog

import (
	"log"
	"time"

	"github.com/sasquad92/server-watch/service"
)

// Watchdog struct
type Watchdog struct {
	attemps         int
	checkInterval   time.Duration
	restartInterval time.Duration
	service         *service.Service
}

// NewWatchdog creates new watchdog struct and initiate its fields
func NewWatchdog(attemps int, checkInterval time.Duration, restartInterval time.Duration, serv *service.Service) *Watchdog {
	watchdog := Watchdog{
		attemps:         attemps,
		checkInterval:   checkInterval,
		restartInterval: restartInterval,
		service:         serv,
	}

	return &watchdog
}

// Watch runs the provided service and next check its status.
// If service is down, function will try to restart it specified
// number of times with specified time interval.
func (w *Watchdog) Watch() {

	//TODO: implement more logic

	for {

		isRunning := w.service.CheckStatus()

		if isRunning == false {

			for i := 0; i < w.attemps; i++ {
				if err := w.service.Start(); err != nil {
					log.Fatal(err.Error())
				}

				// some logic missing

				time.Sleep(w.restartInterval * time.Second)
			}
		}

		time.Sleep(w.checkInterval * time.Second)
	}
}
