package util

import (
	"bailuoxi66/go-loggie/pkg/core/log"
	"time"
)

func AsyncRunWithTimeout(f func(), timeout time.Duration) {
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	done := make(chan struct{})
	go func() {
		now := time.Now()
		f()
		since := time.Since(now)
		if since > timeout {
			log.Info("func run too long(%ds)", since/time.Second)
		}

		done <- struct{}{}
	}()
	for {
		select {
		case <-done:
			return
		case <-timer.C:
			return
		}
	}
}
