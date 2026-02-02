package main

import (
	"log"
	"sync/atomic"
	"time"

	batcher "github.com/Marie20767/go-playground/systems/timedbatcher"
)

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

func main() {
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

	defer close(done)

	select {
	case <-time.After(waitTime + time.Second):
		log.Println("failed to process 1st batch")
	case <-processor.done:
		log.Println("processed batch 1")
	}

	select {
	case <-time.After(waitTime + time.Second):
		log.Println("failed to process 2nd batch")
	case <-processor.done:
		log.Println("processed batch 2")
	}
}
