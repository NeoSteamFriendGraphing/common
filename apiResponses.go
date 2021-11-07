package common

import "time"

type UptimeResponse struct {
	Status string        `json:"status"`
	Uptime time.Duration `json:"uptime"`
}

type LoggingFields struct {
	NodeName string
	NodeDC   string
	LogPaths []string
	NodeIPV4 string
}
