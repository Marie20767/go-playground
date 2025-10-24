package concurrency

import (
	"fmt"
	"time"

	"github.com/Marie20767/go-playground/concurrency/generator"
)

func FanInExample() {

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
