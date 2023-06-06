package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

var ErrNotEnoughGoroutines = errors.New("not enough goroutines")

var ErrNegNumberOfAllowedErrors = errors.New("negative number of allowed errors")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n < 1 {
		return ErrNotEnoughGoroutines
	}

	if m < 0 {
		return ErrNegNumberOfAllowedErrors
	}

	if m > len(tasks) {
		m = len(tasks)
	}

	errLimit := int64(m)
	if m == 0 {
		errLimit += int64(len(tasks)) + 1
	}

	ch := make(chan Task)
	var errCnt int64
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for t := range ch {
				tt := t()
				if tt != nil {
					atomic.AddInt64(&errCnt, 1)
				}
			}
		}()
	}

	for _, t := range tasks {
		if atomic.LoadInt64(&errCnt) >= errLimit {
			break
		}
		ch <- t
	}

	close(ch)
	wg.Wait()

	if atomic.LoadInt64(&errCnt) >= errLimit {
		return ErrErrorsLimitExceeded
	}

	return nil
}
