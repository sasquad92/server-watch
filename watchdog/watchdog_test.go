package watchdog_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/sasquad92/server-watch/configuration"
	"github.com/sasquad92/server-watch/mail"
	"github.com/sasquad92/server-watch/service"
	"github.com/sasquad92/server-watch/watchdog"
)

func GetConfig() *configuration.Configuration {
	config, err := configuration.NewConfigFile("../config.json")

	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}

func GetService(serviceName string) *service.Service {
	serv, err := service.NewService(serviceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	return serv
}

func GetNotifier(config *configuration.Configuration) *mail.Mail {
	notifier, err := mail.NewMail(
		config.NotificationMailTo,
		config.NotificationMailFrom,
		config.MailSmtp,
		config.MailPort,
		config.MailPassword)

	if err != nil {
		log.Fatal(err.Error())
	}

	return notifier
}

func TestNewWatchdog_ServiceNil(t *testing.T) {
	t.Log("Service instance should not be null")

	config := GetConfig()
	notifier := GetNotifier(config)

	_, err := watchdog.NewWatchdog(1, 2, 2, nil, notifier)

	if err != nil {
		expected := fmt.Errorf("Error during creating Watchdog - serv is nil.")

		if err.Error() != expected.Error() {
			t.Errorf("Wrong error ocured. Expected: %s, occured: %s", expected.Error(), err.Error())
		}
	} else {
		t.Error("Error expected.")
	}
}

func TestNewWatchdog_NotifierNil(t *testing.T) {
	t.Log("Notifier instance should not be null")

	config := GetConfig()
	serv := GetService(config.ServiceName)

	_, err := watchdog.NewWatchdog(1, 2, 2, serv, nil)

	if err != nil {
		expected := fmt.Errorf("Error during creating Watchdog - notifier is nil.")

		if err.Error() != expected.Error() {
			t.Errorf("Wrong error ocured. Expected: %s, occured: %s", expected.Error(), err.Error())
		}
	} else {
		t.Error("Error expected.")
	}
}

func TestNewWatchdog_AttempsToSmall(t *testing.T) {
	t.Log("Attemps should be greater than 0")

	config := GetConfig()
	serv := GetService(config.ServiceName)
	notifier := GetNotifier(config)

	_, err := watchdog.NewWatchdog(-1, 2, 2, serv, notifier)

	if err != nil {
		expected := fmt.Errorf("Error during creating Watchdog - attemps less than 0.")

		if err.Error() != expected.Error() {
			t.Errorf("Wrong error ocured. Expected: %s, occured: %s", expected.Error(), err.Error())
		}
	} else {
		t.Error("Error expected.")
	}
}

func TestNewWatchdog_CheckIntervalToSmall(t *testing.T) {
	t.Log("CheckInterval should be greater than 1")

	config := GetConfig()
	serv := GetService(config.ServiceName)
	notifier := GetNotifier(config)

	_, err := watchdog.NewWatchdog(1, 0, 2, serv, notifier)

	if err != nil {
		expected := fmt.Errorf("Error during creating Watchdog - check interval less than 1.")

		if err.Error() != expected.Error() {
			t.Errorf("Wrong error ocured. Expected: %s, occured: %s", expected.Error(), err.Error())
		}
	} else {
		t.Error("Error expected.")
	}
}

func TestNewWatchdog_RestartIntervalToSmall(t *testing.T) {
	t.Log("RestartInterval should be greater than 1")

	config := GetConfig()
	serv := GetService(config.ServiceName)
	notifier := GetNotifier(config)

	_, err := watchdog.NewWatchdog(1, 2, 0, serv, notifier)

	if err != nil {
		expected := fmt.Errorf("Error during creating Watchdog - restart interval less than 1.")

		if err.Error() != expected.Error() {
			t.Errorf("Wrong error ocured. Expected: %s, occured: %s", expected.Error(), err.Error())
		}
	} else {
		t.Error("Error expected.")
	}
}

func TestNewWatchdog_Success(t *testing.T) {
	config := GetConfig()
	serv := GetService(config.ServiceName)
	notifier := GetNotifier(config)

	wd, err := watchdog.NewWatchdog(1, 2, 2, serv, notifier)

	if wd == nil {
		t.Error("Watchdog instance should not be nil.")
	}

	if err != nil {
		t.Error("Error should be nil.")
	}
}
