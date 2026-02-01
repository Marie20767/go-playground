package batcher_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Marie20767/go-playground/systems/batcher"
	"github.com/stretchr/testify/assert"
)

type Job struct {
	ID    int
	Query string
}

type Processor struct {
	done      chan struct{}
	processed atomic.Int32
}

func (p *Processor) Process(jobs []Job) error {
	close(p.done)
	p.processed.Add(int32(len(jobs)))
	return nil
}

func setup(t *testing.T, maxBatchSize int, jobs []Job) (*Processor, *batcher.Batcher[Job]) {
	t.Helper()

	done := make(chan struct{})
	p := &Processor{done: done}
	b := batcher.New(p, maxBatchSize)

	for _, job := range jobs {
		b.Add(job)
	}

	return p, b
}

func TestBatcher(t *testing.T) {
	t.Run("Executes batch immediately when batch size is exceeded", func(t *testing.T) {
		maxBatchSize := 3
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO users (ID, name) VALUES (1, 'Marie')"},
			{ID: 2, Query: "INSERT INTO users (ID, name) VALUES (2, 'Jamie')"},
			{ID: 3, Query: "INSERT INTO users (ID, name) VALUES (3, 'Alfie')"},
		}
		p, _ := setup(t, maxBatchSize, jobs)
		assert.Zero(t, p.processed.Load())

		select {
		case <-time.After(500 * time.Millisecond):
			t.Fatal("Failed to call batch processing function")
		case <-p.done:
			// test passes
		}
	})

	t.Run("Does not execute batch if batch size is not exceeded", func(t *testing.T) {
		maxBatchSize := 2
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO users (ID, name) VALUES (1, 'Marie')"},
		}
		p, _ := setup(t, maxBatchSize, jobs)
		assert.Zero(t, p.processed.Load())

		time.Sleep(500 * time.Millisecond)

		assert.Zero(t, p.processed.Load())
	})

	t.Run("Jobs are completed on close", func(t *testing.T) {
		maxBatchSize := 1
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO users (ID, name) VALUES (1, 'Marie')"},
		}
		p, b := setup(t, maxBatchSize, jobs)
		assert.Zero(t, p.processed.Load())

		ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancelCtx()
		err := b.Close(ctx)
		assert.Nil(t, err)
		assert.Equal(t, int32(len(jobs)), p.processed.Load())
	})

	t.Run("Context cancellation returns error before jobs completion", func(t *testing.T) {
		maxBatchSize := 1
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO users (ID, name) VALUES (1, 'Marie')"},
		}
		p, b := setup(t, maxBatchSize, jobs)
		assert.Zero(t, p.processed.Load())

		ctx, cancelCtx := context.WithTimeout(context.Background(), 1*time.Nanosecond)
		defer cancelCtx()
		err := b.Close(ctx)

		assert.NotNil(t, err)
		assert.Zero(t, p.processed.Load())
	})
}
