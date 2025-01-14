package water

import "fmt"

const deathTime = 480 // 8 minutes
type Status struct {
	Temperature               float64
	LastSuccessfulPollSeconds int64 // in seconds
}

func (s Status) String() string {
	var statusStr = ""
	if s.LastSuccessfulPollSeconds > deathTime || s.Temperature == 0 {
		statusStr = "Server is 100% dead"
	} else {
		statusStr = fmt.Sprintf("%d seconds ago", s.LastSuccessfulPollSeconds)
	}
	return fmt.Sprintf("Water tank temperature is %.2f (%s)", s.Temperature, statusStr)
}
