package internal

import (
	"net/http"
	"strings"
	"time"
)

// / Actual HTTP Call
func HttpCall(url string, requestNumber int64, logChannel chan<- RequestLog) {
	startTime := time.Now()

	// Ensure URL starts with http or https
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	// Perform the HTTP GET request
	res, err := http.Get(url)
	if err != nil {
		logChannel <- RequestLog{
			StatusCode:       0,
			RequestNumber:    requestNumber,
			RequestStartTime: startTime,
			RequestEndTime:   time.Now(),
			RequestDuration:  time.Since(startTime),
		}
		return
	}
	defer res.Body.Close()

	logChannel <- RequestLog{
		StatusCode:       res.StatusCode,
		RequestNumber:    requestNumber,
		RequestStartTime: startTime,
		RequestEndTime:   time.Now(),
		RequestDuration:  time.Since(startTime),
	}
}
