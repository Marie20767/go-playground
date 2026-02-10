package main

import (
	"context"
	"log"
	"sync/atomic"
	"time"

	batcher "github.com/Marie20767/go-playground/systems/combinedbatcher"
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

	return nil
}

func main() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()

	batchSize := 3
	waitTime := 1 * time.Second
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

	processor := &Processor{}
	batcher := batcher.New(processor, batchSize, waitTime)

	for _, job := range jobs {
		batcher.Add(job)
	}

	batcher.Close(ctx)
}
