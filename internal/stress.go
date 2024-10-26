package internal

import (
	"fmt"
	"sync"
)

type StressTest struct {
	URL          string
	RequestCount int64
	Concurrency  int64
}

func NewStressTest(url string, requestCount, concurrency int64) (*StressTest, error) {
	st := &StressTest{
		URL:          url,
		RequestCount: requestCount,
		Concurrency:  concurrency,
	}

	err := st.validate()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return st, nil
}

// TODO: conclude run
func (st *StressTest) Run() {
	fl := NewFinalLog(st.URL, 0, 0, 0.0, st.RequestCount)
	wg := sync.WaitGroup{}
	logCN := make(chan RequestLog, st.RequestCount)
	var mu sync.Mutex // To safely increment reqN
	reqPerWorker := st.RequestCount / st.Concurrency
	remainingReq := st.RequestCount % st.Concurrency
	reqN := int64(0)

	// Goroutine to handle logging and display the loading indicator
	go func() {
		loadingChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0
		for reqLog := range logCN {
			// Display loading spinner and request log info on the same line
			fmt.Printf("\r%s Processing request %d, status code: %d, duration: %v",
				loadingChars[i%len(loadingChars)], reqLog.RequestNumber, reqLog.StatusCode, reqLog.RequestDuration)
			i++

			// Update the final log with statistics
			fl.TimeTaken += reqLog.RequestDuration
			fl.RequestsMade++
			if reqLog.StatusCode >= 200 && reqLog.StatusCode < 300 {
				fl.SuccessReqCount++
			} else {
				fl.AddErrorRequest(reqLog.StatusCode)
			}
		}
		fmt.Println() // Newline after all logs processed
	}()

	// Start worker goroutines
	for i := int64(0); i < st.Concurrency; i++ {
		wg.Add(1)
		go func(workerID int64) {
			defer wg.Done()
			reqForThisWorker := reqPerWorker
			if workerID < remainingReq {
				reqForThisWorker++
			}

			for j := int64(0); j < reqForThisWorker; j++ {
				mu.Lock()
				currentReq := reqN
				reqN++
				mu.Unlock()

				// Perform the HTTP call and send the result to logCN
				HttpCall(st.URL, currentReq, logCN)
			}
		}(i)
	}

	// Wait for all workers to finish, then close the log channel
	go func() {
		wg.Wait()
		close(logCN)
	}()

	// Final results output
	fmt.Println("Waiting for all goroutines to finish...")
	wg.Wait()
	fmt.Println("------------------------------------------")
	fmt.Println("Stress test completed")
	fl.printLog()
	fmt.Println("------------------------------------------")
}

func (st *StressTest) validate() error {
	if st.URL == "" {
		return fmt.Errorf("URL is required")
	}
	if st.RequestCount == 0 {
		return fmt.Errorf("request Count is required")
	}
	if st.Concurrency == 0 {
		return fmt.Errorf("roncurrency is required")
	}
	return nil
}
