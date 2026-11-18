package internal

import (
	"net/netip"
	"time"
)

type Event struct {
	Time       time.Time  `json:"time"`
	IP         netip.Addr `json:"ip"`
	Method     string     `json:"method"`
	Path       string     `json:"path"`
	StatusCode int        `json:"status_code"`
	Bytes      int64      `json:"bytes"`
	LineNumber int        `json:"line_number"`
}
