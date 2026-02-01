package batcher

import (
	"context"
	"fmt"
	"sync"
)

type BatchProcessor[T any] interface {
	Process(jobs []T) error
}

type Batcher[T any] struct {
	processor    BatchProcessor[T]
	jobs         []T
	mu           sync.Mutex
	wg           sync.WaitGroup
	maxBatchSize int
}

func New[T any](batchProcessor BatchProcessor[T], maxBatchSize int) *Batcher[T] {
	return &Batcher[T]{
		processor:    batchProcessor,
		maxBatchSize: maxBatchSize,
		jobs:         []T{},
	}
}

func (b *Batcher[T]) Add(j T) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.jobs = append(b.jobs, j)

	if len(b.jobs) >= b.maxBatchSize {
		b.execute()
	}
}

func (b *Batcher[T]) Execute() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.execute()
}

func (b *Batcher[T]) execute() {
	if len(b.jobs) == 0 {
		return
	}

	currentJobs := make([]T, len(b.jobs))
	copy(currentJobs, b.jobs)
	b.jobs = []T{}

	b.wg.Go(func() {
		if err := b.processor.Process(currentJobs); err != nil {
			fmt.Println("error processing batch", err)
		}
	})
}

func (b *Batcher[T]) Wait() {
	b.wg.Wait()
}

func (b *Batcher[T]) Close(ctx context.Context) error {
	done := make(chan struct{})

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
