package batcher_test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	batcher "github.com/Marie20767/go-playground/systems/combinedbatcher"
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
	fmt.Printf("processing %d jobs", len(jobs))
	close(p.done)
	p.processed.Add(int32(len(jobs)))
	return nil
}

func TestCombinedBatcher(t *testing.T) {
	t.Run("Executes batch immediately", func(t *testing.T) {
		maxBatchSize := 10
		waitTime := 200 * time.Millisecond
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO users (ID, name) VALUES (1, 'Marie')"},
			{ID: 2, Query: "INSERT INTO users (ID, name) VALUES (2, 'Jamie')"},
			{ID: 3, Query: "INSERT INTO users (ID, name) VALUES (3, 'Alfie')"},
		}

		done := make(chan struct{})
		processor := &Processor{done: done}
		batcher := batcher.New(processor, maxBatchSize, waitTime)

		assert.Zero(t, processor.processed.Load())

		for _, job := range jobs {
			batcher.Add(job)
		}

		select {
		case <-time.After(300 * time.Millisecond):
			t.Fatalf("Failed to execute batch")
		case <-done:
			assert.Equal(t, int32(len(jobs)), processor.processed.Load())
		}
	})

	t.Run("Executes batch when maxBatchSize limit is reached", func(t *testing.T) {
		maxBatchSize := 3
		waitTime := 1 * time.Second
		jobs := []Job{
			{ID: 1, Query: "INSERT INTO users (ID, name) VALUES (1, 'Marie')"},
			{ID: 2, Query: "INSERT INTO users (ID, name) VALUES (2, 'Jamie')"},
			{ID: 3, Query: "INSERT INTO users (ID, name) VALUES (3, 'Alfie')"},
		}

		done := make(chan struct{})
		processor := &Processor{done: done}
		batcher := batcher.New(processor, maxBatchSize, waitTime)

		assert.Zero(t, processor.processed.Load())

		for _, job := range jobs {
			batcher.Add(job)
		}

		select {
		case <-time.After(waitTime + 200*time.Millisecond):
			t.Fatalf("Failed to execute batch")
		case <-done:
			assert.Equal(t, int32(len(jobs)), processor.processed.Load())
		}
	})

	t.Run("Does not add any new jobs on close", func(t *testing.T) {

	})

	t.Run("Executes remaining batch on close", func(t *testing.T) {

	})
}
