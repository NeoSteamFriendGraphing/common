package common

import "time"

// UptimeResponse is the standard response
// for any service's /status endpoint
type UptimeResponse struct {
	Status string        `json:"status"`
	Uptime time.Duration `json:"uptime"`
}
