package main

import (
	"context"
	"fmt"
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

	return nil
}

func main() {
	batchSize := 10
	waitTime := 50 * time.Millisecond
	jobs := []Job{
		{ID: 1, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
		{ID: 2, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
		{ID: 3, Query: "INSERT INTO USERS (name) VALUES ('Marie')"},
	}
	processor := &Processor{}
	batcher := batcher.New(processor, batchSize, waitTime)

	for _, job := range jobs {
		batcher.Add(job)
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancelCtx()

	err := batcher.Close(ctx)
	fmt.Println(">>> err: ", err)
}
