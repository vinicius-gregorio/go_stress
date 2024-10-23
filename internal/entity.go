package internal

import (
	"time"
)

type RequestLog struct {
	StatusCode       int
	RequestNumber    int64
	RequestStartTime time.Time
	RequestEndTime   time.Time
	RequestDuration  time.Duration
}
