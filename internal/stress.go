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

func (st *StressTest) Run() {
	fl := NewFinalLog(st.URL, 0, 0, 0.0, st.RequestCount)
	wg := sync.WaitGroup{}
	logCN := make(chan RequestLog, st.RequestCount)
	var mu sync.Mutex
	reqPerWorker := st.RequestCount / st.Concurrency
	remainingReq := st.RequestCount % st.Concurrency
	reqN := int64(0)

	go func() {
		loadingChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0
		for reqLog := range logCN {

			fmt.Printf("\r%s Processing request %d, status code: %d, duration: %v",
				loadingChars[i%len(loadingChars)], reqLog.RequestNumber, reqLog.StatusCode, reqLog.RequestDuration)
			i++

			fl.TimeTaken += reqLog.RequestDuration
			fl.RequestsMade++
			if reqLog.StatusCode >= 200 && reqLog.StatusCode < 300 {
				fl.SuccessReqCount++
			} else {
				fl.AddErrorRequest(reqLog.StatusCode)
			}
		}
		fmt.Println()
	}()

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

				HttpCall(st.URL, currentReq, logCN)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(logCN)
	}()

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
