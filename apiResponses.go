package common

import "time"

// UptimeResponse is the standard response
// for any service's /status endpoint
type UptimeResponse struct {
	Status string        `json:"status"`
	Uptime time.Duration `json:"uptime"`
}

// BasicAPIResponse is the basic API response
// for endpoints that just need a response
// rather than data to be returned
type BasicAPIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
