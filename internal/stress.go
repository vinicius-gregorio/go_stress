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
	wg := sync.WaitGroup{}
	logCN := make(chan RequestLog, st.Concurrency*st.RequestCount)
	defer close(logCN)

	go func() {
		for log := range logCN {
			fmt.Printf("Request %d received status code: %d, duration: %v\n", log.RequestNumber, log.StatusCode, log.RequestDuration)
		}
	}()

	for i := int64(0); i < st.Concurrency; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()
			for j := int64(0); j <= st.RequestCount; j++ {
				HttpCall(st.URL, j, logCN)
			}
		}(i)
	}
	fmt.Println("Waiting for all goroutines to finish...")
	wg.Wait()
	fmt.Println("Stress test completed")

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
