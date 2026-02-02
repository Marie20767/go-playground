package batcher

import (
	"log"
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
}

func New[J any](processor BatchProcessor[J], batchSize int, waitTime time.Duration) *Batcher[J] {
	b := &Batcher[J]{
		processor: processor,
		jobs:      []J{},
		batchSize: batchSize,
		ticker:    time.NewTicker(waitTime),
		waitTime:  waitTime,
	}

	go b.run()

	return b
}

func (b *Batcher[J]) Add(job J) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.jobs = append(b.jobs, job)
}

func (b *Batcher[J]) run() {
	for range b.ticker.C {
		b.execute()
	}
}

func (b *Batcher[J]) execute() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.jobs) == 0 {
		return
	}

	batchSize := b.batchSize
	if len(b.jobs) <= b.batchSize {
		batchSize = (len(b.jobs))
	}

	batch := []J{}
	batch = append(batch, b.jobs[:batchSize]...)
	b.jobs = b.jobs[batchSize:]

	go func() {
		if err := b.processor.Process(batch); err != nil {
			log.Printf("error processing batch %v", err)
		}
	}()
}
