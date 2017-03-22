package watchdog

import (
	"fmt"
	"log"
	"time"

	"github.com/sasquad92/server-watch/mail"
	"github.com/sasquad92/server-watch/service"
)

// Watchdog struct
type Watchdog struct {
	attemps         int
	checkInterval   time.Duration
	restartInterval time.Duration
	service         *service.Service
	notifier        *mail.Mail
}

// NewWatchdog creates new watchdog struct and initiate its fields
func NewWatchdog(attemps int, checkInterval time.Duration, restartInterval time.Duration, serv *service.Service, notifier *mail.Mail) *Watchdog {
	watchdog := Watchdog{
		attemps:         attemps,
		checkInterval:   checkInterval,
		restartInterval: restartInterval,
		service:         serv,
		notifier:        notifier,
	}

	return &watchdog
}

// Watch runs the provided service and next check its status.
// If service is down, function will try to restart it specified
// number of times with specified time interval.
func (w *Watchdog) Watch() {

	for {

		isRunning := w.service.CheckStatus()

		if isRunning == false {
			w.sendNotificationDown()

			for i := 0; i < w.attemps; i++ {
				if err := w.service.Start(); err != nil {
					log.Fatal(err.Error())
				}

				if isRunning = w.service.CheckStatus(); isRunning {
					w.sendNotificationRestarted(i)
					break
				}

				time.Sleep(w.restartInterval * time.Second)
			}

			if !isRunning {
				w.sendNotificationNotRestarted()
				log.Fatal("Service not responding!")
			}
		}

		time.Sleep(w.checkInterval * time.Second)
	}
}

// sendNotificationDown sends an email at provided address
// with an information that service is down
func (w *Watchdog) sendNotificationDown() {
	serviceName := w.service.GetName()
	subject := fmt.Sprintf("Service %s is down", serviceName)
	body := fmt.Sprintf("Service %s is down. Watchdog will try to restart it.", serviceName)
	err := w.notifier.Send(subject, body)

	if err != nil {
		log.Fatal(err.Error())
	}
}

// sendNotificationRestarted sends an email at provided address
// with an information that service has been restarted
func (w *Watchdog) sendNotificationRestarted(attemps int) {
	serviceName := w.service.GetName()
	subject := fmt.Sprintf("Service %s is now running", serviceName)
	body := fmt.Sprintf("Service %s is running afer %d attemps of restart.", serviceName, attemps)
	err := w.notifier.Send(subject, body)

	if err != nil {
		log.Fatal(err.Error())
	}
}

// sendNotificationNotRestarted sends an email at provided address
// with an information that service has not been restarted
func (w *Watchdog) sendNotificationNotRestarted() {
	serviceName := w.service.GetName()
	subject := fmt.Sprintf("Service %s is down", serviceName)
	body := fmt.Sprintf("Service %s is down. Watchdog tried to restart it %d times.", serviceName, w.attemps)
	err := w.notifier.Send(subject, body)

	if err != nil {
		log.Fatal(err.Error())
	}
}
