package internal

import (
	"fmt"
	"time"
)

type RequestLog struct {
	StatusCode       int
	RequestNumber    int64
	RequestStartTime time.Time
	RequestEndTime   time.Time
	RequestDuration  time.Duration
}

type FinalLog struct {
	URL             string
	RequestsMade    int64
	TimeTaken       time.Duration
	SuccessReqCount int64
	TotalReqCount   int64
	ErrorRequests   map[int]int64
}

func NewFinalLog(url string, requestsMade int64, timeTaken time.Duration, successRate int64, totalReqCount int64) *FinalLog {
	return &FinalLog{
		URL:             url,
		RequestsMade:    requestsMade,
		TimeTaken:       timeTaken,
		SuccessReqCount: successRate,
		TotalReqCount:   totalReqCount,
	}
}

func (fl *FinalLog) AddErrorRequest(statusCode int) {
	if fl.ErrorRequests == nil {
		fl.ErrorRequests = make(map[int]int64)
	}
	fl.ErrorRequests[statusCode]++
}
func (fl *FinalLog) printLog() {
	fmt.Printf("Number of requests made: %d\n", fl.RequestsMade)
	fmt.Printf("Time taken: %v\n", fl.TimeTaken)
	fmt.Printf(
		"Success rate: %d %% \n",
		(fl.SuccessReqCount/fl.TotalReqCount)*100,
	)

	/// This prints the error request count
	/// Should arrange the status codes that has the same code, for example:
	/// If there are 3 requests with status code 404, and 2 with error 500:
	/// it should print: "Error request count: 404: 3, 500: 2"
	fmt.Println("Error request count:")
	for statusCode, count := range fl.ErrorRequests {
		fmt.Printf("%d: %d\n", statusCode, count)
	}

}
