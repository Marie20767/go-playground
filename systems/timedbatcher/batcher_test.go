package batcher_test

import (
	"context"
	"errors"
	"log"
	"sync/atomic"
	"testing"
	"time"

	batcher "github.com/Marie20767/go-playground/systems/timedbatcher"
	"github.com/stretchr/testify/assert"
)

var ErrBatchFailed = errors.New("batch failed")

type Job struct {
	ID    int
	Query string
}

type Processor struct {
	processed atomic.Int32
	done      chan (struct{})
}

type ProcessCloseJobs struct {
	processed atomic.Int32
}

type ProcessErr struct {
	done      chan (struct{})
	processed atomic.Int32
}

func (p *Processor) Process(jobs []Job) error {
	log.Printf("jobs to process: %v", len(jobs))
	p.processed.Add(int32(len(jobs)))
	p.done <- struct{}{}
	return nil
}

func (p *ProcessCloseJobs) Process(jobs []Job) error {
	log.Printf("jobs to process: %v", len(jobs))
	p.processed.Add(int32(len(jobs)))
	return nil
}

func (p *ProcessErr) Process(jobs []Job) error {
	return ErrBatchFailed
}

func TestTimedBatcher(t *testing.T) {
	t.Run("Processes batches of up to 10 jobs every 500ms", func(t *testing.T) {
		batchSize := 10
		waitTime := 500 * time.Millisecond
		done := make(chan struct{})
		processor := Processor{done: done}
		batcher := batcher.New(&processor, batchSize, waitTime)
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 2, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 3, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 4, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 5, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 6, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 7, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 8, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 9, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 10, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 11, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 12, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
		}

		for _, job := range jobs {
			batcher.Add(job)
		}

		assert.Zero(t, processor.processed.Load())
		defer close(done)

		select {
		case <-time.After(waitTime + time.Second):
			t.Fatalf("failed to process first batch")
		case <-processor.done:
			log.Println("processed batch 1")
			assert.Equal(t, int32(batchSize), processor.processed.Load())
		}

		select {
		case <-time.After(waitTime + time.Second):
			t.Fatalf("failed to process second batch")
		case <-processor.done:
			log.Println("processed batch 2")
			assert.Equal(t, int32(len(jobs)), processor.processed.Load())
		}
	})

	t.Run("Context cancellation returns error before jobs completion", func(t *testing.T) {
		batchSize := 10
		waitTime := 500 * time.Millisecond
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 2, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 3, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
		}
		processor := &Processor{}
		batcher := batcher.New(processor, batchSize, waitTime)
		assert.Zero(t, processor.processed.Load())

		for _, job := range jobs {
			batcher.Add(job)
		}

		ctx, cancelCtx := context.WithTimeout(t.Context(), 1*time.Nanosecond)
		defer cancelCtx()
		err := batcher.Close(ctx)
		assert.NotNil(t, err)
		assert.Zero(t, processor.processed.Load())
	})

	t.Run("Jobs are completed on close", func(t *testing.T) {
		batchSize := 10
		waitTime := 50 * time.Millisecond
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 2, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 3, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
		}
		processor := &ProcessCloseJobs{}
		batcher := batcher.New(processor, batchSize, waitTime)
		assert.Zero(t, processor.processed.Load())

		for _, job := range jobs {
			batcher.Add(job)
		}

		ctx, cancelCtx := context.WithTimeout(t.Context(), 500*time.Millisecond)
		defer cancelCtx()

		err := batcher.Close(ctx)
		assert.Nil(t, err)
		assert.Equal(t, int32(len(jobs)), processor.processed.Load())
	})

	t.Run("Returns batch error with failed jobs", func(t *testing.T) {
		batchSize := 10
		waitTime := 50 * time.Millisecond
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 2, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
			{ID: 3, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
		}
		processor := &ProcessErr{}
		batcher := batcher.New(processor, batchSize, waitTime)
		assert.Zero(t, processor.processed.Load())

		for _, job := range jobs {
			batcher.Add(job)
		}

		done := make(chan struct{})

		go func() {
			defer close(done)
			for failure := range batcher.Error() {
				assert.Equal(t, len(jobs), len(failure.Jobs))
				assert.Equal(t, failure.Err, ErrBatchFailed)
			}
		}()

		ctx, cancelCtx := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancelCtx()
		batcher.Close(ctx)

		select {
		case <-time.After(500 * time.Millisecond):
			t.Fatal("failed to read failed batch error")
		case <-done:
			// test passes
		}
	})
}
