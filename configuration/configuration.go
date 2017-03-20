package configuration

import (
	"encoding/json"
	"os"
	"time"
)

// Configuration struct represents config file
type Configuration struct {
	Attemps          int
	RestartInterval  time.Duration
	CheckInterval    time.Duration
	NotificationMail string
	ServiceName      string
}

// NewConfigFile creates new Configuration struct based on json config file.
// If successful method returns pointer to Configuration struct; if not:
// the error will be returned.
func NewConfigFile(filename string) (*Configuration, error) {
	var config Configuration
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
