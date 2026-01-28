package batcher

import "fmt"

// TODO: make this generic
type Job struct {
	ID    int
	Query string
}

type Batcher struct {
	processBatch func(jobs []Job) error
	jobs         []Job
}

func New(processBatch func(jobs []Job) error) *Batcher {
	return &Batcher{
		processBatch: processBatch,
		jobs:         []Job{},
	}
}

func (b *Batcher) Add(j Job) {
	b.jobs = append(b.jobs, j)
}

func (b *Batcher) Execute() {
	if err := b.processBatch(b.jobs); err != nil {
		fmt.Println("error processing batch", err)
	}

}

// TODO: handle concurrency
