package service

import "os/exec"
import "fmt"

// Service struct represents service which will be watched
type Service struct {
	name string
}

const (
	SERVICE = "service"
	START   = "start"
	STATUS  = "status"
)

// NewService creates new service instance and return reference to service object
func NewService(name string) (*Service, error) {
	
	if !(len(name) > 0) {
		return nil, fmt.Errorf("Service name is incorrect.")
	}

	service := Service{
		name: name,
	}

	return &service, nil
}

// Start starts service, returns error - if occured. Otherwise
// error is nil
func (s *Service) Start() error {

	_, err := exec.Command(SERVICE, s.name, START).Output()

	if err != nil {
		return fmt.Errorf("Error during starting %s service.", s.name)
	}

	return nil
}

// CheckStatus checks if service is running
func (s *Service) CheckStatus() bool {
	_, err := exec.Command(SERVICE, s.name, STATUS).Output()

	if err != nil {
		return false
	}

	return true
}

// GetName returns service name
func (s *Service) GetName() string {
	return s.name
}
