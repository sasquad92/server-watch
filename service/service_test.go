package service_test

import (
	"fmt"
	"testing"

	"github.com/sasquad92/server-watch/service"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestNewService_IncorrectName(t *testing.T) {
	t.Log("Service name should be longer than 0 chars")
	service, err := service.NewService("")

	if service != nil {
		t.Error("Service object should be nil", ballotX)
	} else {
		t.Log("Service object should be nil", checkMark)
	}

	if err != nil {
		t.Log("Error object should not be nil", checkMark)
	} else {
		t.Error("Error object should not be nil", ballotX)
	}
}

func TestStart_ServiceDoesntExist(t *testing.T) {
	t.Log("Service should exist")

	service, err := service.NewService("someWirdService")

	if err != nil {
		t.Error("Could not create service instance", ballotX)
	}

	err = service.Start()

	if err != nil {
		expectedErr := fmt.Errorf("Error during starting %s service.", service.GetName())
		if err.Error() != expectedErr.Error() {
			t.Error("Not expected error occurs", ballotX)
		} else {
			t.Log("Correct error occurs:", err.Error(), checkMark)
		}
	}
}

func TestStart_ServiceExist(t *testing.T) {
	t.Log("Service exists")
	t.Log("This test depends on your configuration")
	service, err := service.NewService("bluetooth")

	if err != nil {
		t.Error("Could not create service instance", ballotX)
	}

	service.Start()
}
