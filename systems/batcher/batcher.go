package batcher

import (
	"context"
	"sync"
)

// TODO: timed version

type BatchProcessor[J any] interface {
	Process(jobs []J) error
}

type Batcher[J any] struct {
	processor    BatchProcessor[J]
	jobs         []J
	mu           sync.Mutex
	wg           sync.WaitGroup
	maxBatchSize int
	failures     chan FailedBatch[J]
}

type FailedBatch[J any] struct {
	Jobs []J
	Err  error
}

func New[J any](batchProcessor BatchProcessor[J], maxBatchSize int) *Batcher[J] {
	return &Batcher[J]{
		processor:    batchProcessor,
		maxBatchSize: maxBatchSize,
		jobs:         []J{},
		failures:     make(chan FailedBatch[J], 100),
	}
}

func (b *Batcher[J]) Add(j J) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.jobs = append(b.jobs, j)

	if len(b.jobs) >= b.maxBatchSize {
		b.execute()
	}
}

func (b *Batcher[J]) Execute() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.execute()
}

func (b *Batcher[J]) execute() {
	if len(b.jobs) == 0 {
		return
	}

	currentJobs := make([]J, len(b.jobs))
	copy(currentJobs, b.jobs)
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

func (b *Batcher[J]) Wait() {
	b.wg.Wait()
}

func (b *Batcher[J]) Close(ctx context.Context) error {
	done := make(chan struct{})

	defer close(b.failures)

	go func() {
		b.Wait()
		close(done)
	}()

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
