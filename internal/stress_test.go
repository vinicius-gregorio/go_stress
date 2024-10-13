package internal

import (
	"fmt"
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
	// wg := sync.WaitGroup{}
	fmt.Println("StressTest called")

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
