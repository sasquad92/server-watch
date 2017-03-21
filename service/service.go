package service

import "os/exec"
import "fmt"

// Service struct represents service which will be watched
type Service struct {
	name     string
	path     string
	fullPath string
}

const (
	SERVICE = "service"
	START   = "start"
	STATUS  = "status"
)

// NewService creates new service instance and return reference to service object
func NewService(name string, path string) *Service {
	service := Service{
		name:     name,
		path:     path,
		fullPath: path + name,
	}

	return &service
}

// Start starts service, returns error - if occured. Otherwise
// error is nil
func (s *Service) Start() error {

	_, err := exec.Command(SERVICE, s.fullPath, START).Output()

	if err != nil {
		return fmt.Errorf("Error during starting %s service in %s.", s.name, s.path)
	}

	return nil
}

// CheckStatus checks if service is running
func (s *Service) CheckStatus() bool {
	_, err := exec.Command(SERVICE, s.fullPath, STATUS).Output()

	if err != nil {
		return false
	}

	return true
}
