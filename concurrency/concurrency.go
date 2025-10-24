package concurrency

import (
	"fmt"

	"github.com/Marie20767/go-playground/concurrency/generator"
)

func FanInExample() {

}

func TickerGeneratorExample() {

}

func TickerWithDoneChannel() {

}

func FibGeneratorExample() {

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
