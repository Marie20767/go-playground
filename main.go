package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Marie20767/go-playground/systems/batcher"
)

type Job struct {
	ID    int
	Query string
}

type Processor struct{}
func (p *Processor) Process(jobs []Job) error {
	for i, job := range jobs {
		time.Sleep(100 * time.Millisecond)
		log := fmt.Sprintf("processing job %d with id %d", i, job.ID)
		fmt.Println(log)
	}

	return nil
}

func main() {
	maxBatchSize := 3
	jobs := []Job{
		{ID: 1, Query: "INSERT INTO users (id, name) VALUES (1, 'Marie')"},
		{ID: 2, Query: "INSERT INTO users (id, name) VALUES (2, 'Jamie')"},
		{ID: 3, Query: "INSERT INTO users (id, name) VALUES (3, 'Alfie')"},
	}
	
	b := batcher.New(&Processor{}, maxBatchSize)

	for _, job := range jobs {
		b.Add(job)
	}

	b.Execute()
	ctx, cancelCtx := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancelCtx()
	err := b.Close(ctx)
	if err != nil {
		fmt.Println("jobs unfinished")
	}
}
