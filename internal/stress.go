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
	logCN := make(chan RequestLog, st.Concurrency*st.RequestCount)
	defer close(logCN)

	go func() {
		for log := range logCN {
			fmt.Printf("Request %d received status code: %d, duration: %v\n", log.RequestNumber, log.StatusCode, log.RequestDuration)
			/// This Calculates the success rate and final log
			fl.TimeTaken += log.RequestDuration
			fl.RequestsMade++
			if log.StatusCode >= 200 && log.StatusCode < 300 {
				fl.SuccessReqCount++
			} else {
				fl.AddErrorRequest(log.StatusCode)
			}

		}
	}()

	reqPerWorker := st.RequestCount / st.Concurrency
	remainingReq := st.RequestCount % st.Concurrency

	reqN := int64(0)

	for i := int64(0); i < st.Concurrency; i++ {
		wg.Add(1)
		go func(workerID int64) {
			defer wg.Done()

			reqForThisWorker := reqPerWorker
			if workerID < remainingReq {
				reqForThisWorker++
			}

			for j := int64(0); j <= reqForThisWorker; j++ {
				HttpCall(st.URL, reqN, logCN)
				reqN++
			}
		}(i)
	}
	fmt.Println("Waiting for all goroutines to finish...")
	wg.Wait()
	fmt.Println("Stress test completed")
	fl.printLog()

}

func (st *StressTest) validate() error {
	if st.URL == "" {
		return fmt.Errorf("URL is required")
	}
	if st.RequestCount == 0 {
		return fmt.Errorf("Request Count is required")
	}
	if st.Concurrency == 0 {
		return fmt.Errorf("Concurrency is required")
	}
	return nil
}
