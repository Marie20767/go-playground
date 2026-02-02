package batcher_test

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Marie20767/go-playground/systems/batcher"
	"github.com/stretchr/testify/assert"
)

var ErrProcessing = errors.New("failed to process batch")

type Job struct {
	ID    int
	Query string
}

type Processor struct {
	done      chan struct{}
	processed atomic.Int32
}

func (processor *Processor) Process(jobs []Job) error {
	close(processor.done)
	processor.processed.Add(int32(len(jobs)))
	return nil
}

type ErrorProcessor struct{}

func (processor *ErrorProcessor) Process(jobs []Job) error {
	return ErrProcessing
}

func setup(t *testing.T, maxBatchSize int, jobs []Job) (*Processor, *batcher.Batcher[Job]) {
	t.Helper()

	done := make(chan struct{})
	processor := &Processor{done: done}
	batcher := batcher.New(processor, maxBatchSize)

	for _, job := range jobs {
		batcher.Add(job)
	}

	return processor, batcher
}

func TestBatcher(t *testing.T) {
	t.Run("Executes batch immediately when batch size is exceeded", func(t *testing.T) {
		maxBatchSize := 3
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO users (ID, name) VALUES (1, 'Marie')"},
			{ID: 2, Query: "INSERT INTO users (ID, name) VALUES (2, 'Jamie')"},
			{ID: 3, Query: "INSERT INTO users (ID, name) VALUES (3, 'Alfie')"},
		}
		processor, _ := setup(t, maxBatchSize, jobs)
		assert.Zero(t, processor.processed.Load())

		select {
		case <-time.After(500 * time.Millisecond):
			t.Fatal("Failed to call batch processing function")
		case <-processor.done:
			// test passes
		}
	})

	t.Run("Does not execute batch if batch size is not exceeded", func(t *testing.T) {
		maxBatchSize := 2
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO users (ID, name) VALUES (1, 'Marie')"},
		}
		processor, _ := setup(t, maxBatchSize, jobs)
		assert.Zero(t, processor.processed.Load())

		time.Sleep(500 * time.Millisecond)

		assert.Zero(t, processor.processed.Load())
	})

	t.Run("Jobs are completed on close", func(t *testing.T) {
		maxBatchSize := 1
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO users (ID, name) VALUES (1, 'Marie')"},
		}
		processor, batcher := setup(t, maxBatchSize, jobs)
		assert.Zero(t, processor.processed.Load())

		ctx, cancelCtx := context.WithTimeout(t.Context(), 100*time.Millisecond)
		defer cancelCtx()
		err := batcher.Close(ctx)
		assert.Nil(t, err)
		assert.Equal(t, int32(len(jobs)), processor.processed.Load())
	})

	t.Run("Context cancellation returns error before jobs completion", func(t *testing.T) {
		maxBatchSize := 1
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO users (ID, name) VALUES (1, 'Marie')"},
		}
		processor, batcher := setup(t, maxBatchSize, jobs)
		assert.Zero(t, processor.processed.Load())

		ctx, cancelCtx := context.WithTimeout(t.Context(), 1*time.Nanosecond)
		defer cancelCtx()
		err := batcher.Close(ctx)

		assert.NotNil(t, err)
		assert.Zero(t, processor.processed.Load())
	})

	t.Run("Returns batch error with failed jobs", func(t *testing.T) {
		maxBatchSize := 3
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO users (ID, name) VALUES (1, 'Marie')"},
			{ID: 2, Query: "INSERT INTO users (ID, name) VALUES (2, 'Jamie')"},
			{ID: 3, Query: "INSERT INTO users (ID, name) VALUES (3, 'Alfie')"},
		}

		processor := &ErrorProcessor{}
		batcher := batcher.New(processor, maxBatchSize)
		for _, job := range jobs {
			batcher.Add(job)
		}

		done := make(chan struct{})
		go func() {
			defer close(done)
			for failure := range batcher.Error() {
				assert.Equal(t, len(jobs), len(failure.Jobs))
				assert.Equal(t, ErrProcessing, failure.Err)
			}
		}()

		batcher.Close(t.Context())

		select {
		case <-time.After(500 * time.Millisecond):
			t.Fatal("Failed to read failed batch error")
		case <-done:
			// test passes
		}
	})
}
