package main

import (
	"fmt"

	"github.com/Marie20767/go-playground/systems/batcher"
)

func main() {
	// datastructures.BSTSearchExamples()
	// datastructures.MatrixExamples()
	// concurrency.ConcurrencyExamples()

	// combinationsum.CombinationSum([]int{1, 2, 3}, 3)
	b := batcher.New(func(jobs []batcher.Job) error {
		for i, job := range jobs {
			log := fmt.Sprintf("processing job %d with id %d", i, job.ID)
			fmt.Println(log)
		}

		return nil
	})

	b.Add(batcher.Job{ID: 1, Query: "INSERT INTO users (id, name) VALUES (1, 'Marie)"})
	b.Add(batcher.Job{ID: 2, Query: "INSERT INTO users (id, name) VALUES (2, 'Jamie)"})
	b.Add(batcher.Job{ID: 3, Query: "INSERT INTO users (id, name) VALUES (3, 'Tom)"})
	b.Execute()
}
