package watchdog

import "time"

// Watchdog struct
type Watchdog struct {
	Attemps         int
	CheckInterval   time.Duration
	RestartInterval time.Duration
}

// NewWatchdog creates new watchdog struct and initiate its fields
func NewWatchdog(attemps int, checkInterval time.Duration, restartInterval time.Duration) *Watchdog {
	watchdog := Watchdog{
		Attemps:         attemps,
		CheckInterval:   checkInterval,
		RestartInterval: restartInterval,
	}

	return &watchdog
}

// just for test
func (w *Watchdog) GetAttemps() int {
	return w.Attemps
}
