package generator

import (
	"time"
)

// takes an int (num), returns a read only channel
// writes integers into the channel from 0 to num
func Integer(num int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i <= num; i++ {
			ch <- i
		}
	}()

	return ch
}

// takes a read only channel of ints, returns a read only channel of ints
// reads from the input channel and writes double the val from input channel to another channel
func Double(initialCh <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for num := range initialCh {
			ch <- num * 2
		}
	}()

	return ch
}

// takes any number of string args, returns a read only channel of strings
func Strings(s ...string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for _, word := range s {
			ch <- word
		}
	}()

	return ch
}

// returns a read only channel that sends the fibonacci numbers into the channel
func Fibonacci(done <-chan struct{}) <-chan int {
	ch := make(chan int)
	a, b := 0, 1

	go func() {
		defer close(ch)
		for {
			select {
			case ch <- a:
				a, b = b, a+b
			case <-done:
				return
			}
		}
	}()

	return ch
}

// takes a time.Duration and returns a read only channel of time.Time
// sends the current time.Time value into the channel every tick of the given time.Duration
func Ticker(duration time.Duration) <-chan time.Time {
	ticker := time.NewTicker(duration)
	ch := make(chan time.Time)

	go func() {
		defer close(ch)
		defer ticker.Stop()
		for {
			time := <-ticker.C
			ch <- time
		}
	}()

	return ch
}

func TickerWithDone(done <-chan struct{}, duration time.Duration) <-chan time.Time {
	ticker := time.NewTicker(duration)
	ch := make(chan time.Time)

	go func() {
		defer close(ch)
		defer ticker.Stop()
		for {
			select {
			case time := <-ticker.C:
				ch <- time
			case <-done:
				return
			}
		}
	}()

	return ch
}

func Cancellable() {
}
