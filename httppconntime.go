package httppconntime

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

type Probe struct {
	initialWaitTime time.Duration
	maxWaitTime     time.Duration
	logger          *log.Logger
}

func NewProbe(initialTime, maxTime time.Duration, logger *log.Logger) *Probe {
	return &Probe{
		initialWaitTime: initialTime,
		maxWaitTime:     maxTime,
		logger:          logger,
	}
}

func (p *Probe) Do(targetURL string) (time.Duration, error) {
	startWaitTime := p.initialWaitTime
	endWaitTime := p.maxWaitTime
	var lastAttemptWaitTime time.Duration
	for startWaitTime <= endWaitTime {
		median := ((startWaitTime + endWaitTime) / 2).Truncate(time.Second)
		p.logger.Printf("Probe %v - %v", lastAttemptWaitTime, median)
		lastAttemptWaitTime = median

		reused, err := checkReuseConnection(targetURL, median)
		if err != nil {
			return -1, err
		}
		if reused {
			// Connection reused
			startWaitTime = median + time.Second
		} else {
			// Connection discarded
			endWaitTime = median - time.Second
		}
	}

	return lastAttemptWaitTime, nil
}

func checkReuseConnection(targetURL string, waitTime time.Duration) (bool, error) {
	t := &transport{}
	req, _ := http.NewRequest("GET", targetURL, nil)
	trace := &httptrace.ClientTrace{
		GotConn: t.GotConn,
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	client := &http.Client{Transport: t}

	// First Request
	if err := request(client, req); err != nil {
		return false, err
	}

	// Wait
	time.Sleep(waitTime)

	// Second Request
	if err := request(client, req); err != nil {
		return false, err
	}

	// Return reused status
	return t.Reused, nil
}

func request(client *http.Client, req *http.Request) error {
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	_, _ = io.Copy(ioutil.Discard, res.Body)
	_ = res.Body.Close()
	return err
}
