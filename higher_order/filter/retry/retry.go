package main

import (
	"errors"
	"time"
)

func unreliableOperation() error {
	// random failure
	if time.Now().UnixNano()%2 == 0 {
		return errors.New("operation failed")
	}
	return nil
}

func retryWithExponentialBackoff(operation func() error, maxRetries int) error {
	seconds := 1

	for i := 1; i <= maxRetries; i++ {
		err := operation()

		if i == maxRetries {
			return err
		}

		if err == nil {
			return nil
		}

		time.Sleep((time.Duration(seconds)*time.Second) )
		seconds *= 2
	}

	return nil
}

func main() {
	retryWithExponentialBackoff(unreliableOperation, 4)
}

