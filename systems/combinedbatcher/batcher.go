package batcher

import (
	"context"
	"sync"
	"time"
)

type BatchProcessor[J any] interface {
	Process(jobs []J) error
}

type Batcher[J any] struct {
	processor    BatchProcessor[J]
	jobs         []J
	maxBatchSize int
	waitTime     time.Duration
	ticker       *time.Ticker
	mu           sync.Mutex
	wg           sync.WaitGroup
	closed       bool
	done         chan struct{}
	failures     chan FailedBatch[J]
}

type FailedBatch[J any] struct {
	Jobs []J
	Err  error
}

func New[J any](processor BatchProcessor[J], maxBatchSize int, waitTime time.Duration) *Batcher[J] {
	b := &Batcher[J]{
		processor:    processor,
		jobs:         []J{},
		maxBatchSize: maxBatchSize,
		ticker:       time.NewTicker(waitTime),
		waitTime:     waitTime,
		closed:       false,
		failures:     make(chan FailedBatch[J], 100),
		done:         make(chan struct{}),
	}

	b.wg.Go(b.run)

	return b
}

func (b *Batcher[J]) Add(job J) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}

	b.jobs = append(b.jobs, job)

	if len(b.jobs) >= b.maxBatchSize {
		b.execute()
	}
}

func (b *Batcher[J]) run() {
	defer b.ticker.Stop()

	select {
	case <-b.done:
		return
	case <-b.ticker.C:
		b.executeLock()
	}
}

func (b *Batcher[J]) executeLock() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.execute()
}

func (b *Batcher[J]) execute() {
	if len(b.jobs) == 0 {
		return
	}

	currentJobs := b.jobs
	b.jobs = []J{}

	b.wg.Go(func() {
		if err := b.processor.Process(currentJobs); err != nil {
			b.failures <- FailedBatch[J]{
				Jobs: currentJobs,
				Err:  err,
			}
		}
	})
}

func (b *Batcher[J]) wait() {
	b.wg.Wait()
}

func (b *Batcher[J]) Close(ctx context.Context) error {
	close(b.done)
	b.mu.Lock()
	b.closed = true
	b.execute()
	b.mu.Unlock()

	done := make(chan struct{})

	// Note: this would panic if Close() is called again, add channel closure check and return error
	defer close(b.failures)

	go func() {
		b.wait()
		close(done)
	}()

	// Note: currently if context is cancelled the done channel wouldn't close
	// Solution: add context & cancelCtx to batcher, call b.cancelCtx() and listen to ctx.Done() in BatchProcessor.Process

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}
}

func (b *Batcher[J]) Error() <-chan FailedBatch[J] {
	return b.failures
}
