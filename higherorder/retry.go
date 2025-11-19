package higherorder

import (
	"context"
	"errors"
	"time"

	"github.com/cenkalti/backoff/v5"
)

func unreliableOperation() error {
	if time.Now().UnixNano()%2 == 0 {
		return errors.New("operation failed")
	}

	return nil
}

func retryWithExponentialBackoff(maxRetries int, baseDelay time.Duration) error {
	err := unreliableOperation()

	if maxRetries <= 0 || err == nil {
		return err
	}

	time.Sleep(time.Duration(baseDelay))
	return retryWithExponentialBackoff(maxRetries-1, baseDelay*2)
}

func retryWithExponentialBackoff2(ctx context.Context, maxRetries int, baseDelay time.Duration) error {
	operation := func() (struct{}, error) {

		if time.Now().UnixNano()%2 == 0 {
			return struct{}{}, errors.New("operation failed")
		}

		return struct{}{}, nil
	}

	myebo := &backoff.ExponentialBackOff{
		InitialInterval:     baseDelay,
		RandomizationFactor: backoff.DefaultRandomizationFactor,
		Multiplier:          backoff.DefaultMultiplier,
		MaxInterval:         backoff.DefaultMaxInterval,
	}

	_, err := backoff.Retry(ctx, operation, backoff.WithBackOff(myebo), backoff.WithMaxTries(uint(maxRetries)))

	return err
}

func main() {
	ctx := context.Background()

	retryWithExponentialBackoff(4, 3*time.Second)
	retryWithExponentialBackoff2(ctx, 5, 2*time.Second)
}
