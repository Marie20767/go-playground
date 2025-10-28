package concurrency

import (
	"fmt"
	"time"

	"github.com/Marie20767/go-playground/concurrency/fanin"
	"github.com/Marie20767/go-playground/concurrency/generator"
)

func FanInExample() {
	// use generator.Strings to generate 3 different channels of various strings
	// use fanin pattern to combine them into one channnel
	// read from the combined channel and print all the values
	ch1 := generator.Strings("hello 1 channel 1", "hello 2 channel 1")
	ch2 := generator.Strings("hello 1 channel 2", "hello 2 channel 2")
	ch3 := generator.Strings("hello 1 channel 3", "hello 2 channel 3")

	output := fanin.New(ch1, ch2, ch3)
	for s := range output {
		fmt.Println(s)
	}
}

// create a ticker generator that ticks every time.Second
// make it tick for 4 seconds
func TickerGeneratorExample() {
	ch := generator.Ticker(time.Second)

	for range 4 {
		fmt.Println(<-ch)
	}
}

func TickerGeneratorExampleWithDone() {
	done := make(chan struct{})
	ch := generator.TickerWithDone(done, time.Second)

	for range 4 {
		fmt.Println(<-ch)
	}

	close(done)
}

func FibGeneratorExample() {
	done := make(chan struct{})
	ch := generator.Fibonacci(done)

	for range 10 {
		fmt.Println(<-ch)
	}

	close(done)
}

func DoubleGeneratorExample() {
	num := 3
	ch := generator.Integer(num)
	newCh := generator.Double(ch)

	for num := range newCh {
		fmt.Println(num)
	}
}

func IntegerGeneratorExample() {
	num := 4
	ch := generator.Integer(num)

	for num := range ch {
		fmt.Println(num)
	}
}

func StringGeneratorExample() {
	ch := generator.Strings("hello", "my", "name", "is", "something", "or", "other")

	for s := range ch {
		fmt.Println(s)
	}
}

func CancellableExample() {
}
