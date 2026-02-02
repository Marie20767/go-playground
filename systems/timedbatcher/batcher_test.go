package batcher_test

import (
	"log"
	"sync/atomic"
	"testing"
	"time"

	batcher "github.com/Marie20767/go-playground/systems/timedbatcher"
	"github.com/stretchr/testify/assert"
)

// TODO: graceful shutdown
// TODO: error handling

type Job struct {
	ID    int
	Query string
}

type Processor struct {
	processed atomic.Int32
	done      chan (struct{})
}

func (p *Processor) Process(jobs []Job) error {
	log.Printf("jobs to process: %v", len(jobs))
	p.processed.Add(int32(len(jobs)))
	p.done <- struct{}{}
	return nil
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

	t.Run("Jobs are completed on close", func(t *testing.T) {})

	t.Run("Context cancellation returns error before jobs completion", func(t *testing.T) {})

	t.Run("Returns batch error with failed jobs", func(t *testing.T) {})

}
