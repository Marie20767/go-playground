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
	processor BatchProcessor[J]
	jobs      []J
	batchSize int
	waitTime  time.Duration
	ticker    *time.Ticker
	mu        sync.Mutex
	wg        sync.WaitGroup
	closed    bool
	failures  chan FailedBatch[J]
}

type FailedBatch[J any] struct {
	Jobs []J
	Err  error
}

func New[J any](processor BatchProcessor[J], batchSize int, waitTime time.Duration) *Batcher[J] {
	b := &Batcher[J]{
		processor: processor,
		jobs:      []J{},
		batchSize: batchSize,
		ticker:    time.NewTicker(waitTime),
		waitTime:  waitTime,
		closed:    false,
		failures:  make(chan FailedBatch[J], 100),
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
}

func (b *Batcher[J]) run() {
	defer b.ticker.Stop()

	for range b.ticker.C {
		b.execute()
		if b.shouldEndRun() {
			return
		}
	}
}

func (b *Batcher[J]) shouldEndRun() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.closed && len(b.jobs) == 0
}

func (b *Batcher[J]) execute() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.jobs) == 0 {
		return
	}

	batchSize := min(len(b.jobs), b.batchSize)

	batch := []J{}
	batch = append(batch, b.jobs[:batchSize]...)
	b.jobs = b.jobs[batchSize:]

	b.wg.Go(func() {
		if err := b.processor.Process(batch); err != nil {
			b.failures <- FailedBatch[J]{
				Jobs: batch,
				Err:  err,
			}
		}
	})
}

func (b *Batcher[J]) wait() {
	b.wg.Wait()
}

func (b *Batcher[J]) Close(ctx context.Context) error {
	b.mu.Lock()
	b.closed = true
	b.mu.Unlock()
	done := make(chan struct{})

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
