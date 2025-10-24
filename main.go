package main

import "github.com/Marie20767/go-playground/concurrency"

func main() {
	concurrencyExamples()
}

func concurrencyExamples() {
	// ***************************************
	// Fan In Pattern
	// ***************************************
	// concurrency.FanInExample()

	// ***************************************
	// Generator Pattern
	// ***************************************
	// concurrency.IntegerGeneratorExample()
	// concurrency.DoubleGeneratorExample()
	// concurrency.StringGeneratorExample()
	// concurrency.TickerGeneratorExample()
	concurrency.TickerGeneratorExampleWithDone()
	// concurrency.FibGeneratorExample()

	// ***************************************
	// Misc
	// ***************************************
	// concurrency.TickerWithDoneChannel()
	// concurrency.CancellableExample()
}
