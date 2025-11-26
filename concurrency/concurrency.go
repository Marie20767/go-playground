package concurrency

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/Marie20767/go-playground/concurrency/fanin"
	"github.com/Marie20767/go-playground/concurrency/firstresponse"
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

func FirstResponseExample() {
	resp := firstresponse.Fetch(
		fmt.Sprintf("https://swapi.dev/api/people/%d", rand.Intn(50)),
		"https://catfact.ninja/fact",
		"https://official-joke-api.appspot.com/random_joke",
		"https://api.breakingbadquotes.xyz/v1/quotes",
		"https://api.adviceslip.com/advice",
	)

	data, err := getResponseData(resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("first returned data:")
	fmt.Println(data)
}

func AllResponsesExample() {
	ch := firstresponse.FetchAll(
		fmt.Sprintf("https://swapi.dev/api/people/%d", rand.Intn(50)),
		"https://catfact.ninja/fact",
		"https://official-joke-api.appspot.com/random_joke",
		"https://api.breakingbadquotes.xyz/v1/quotes",
		"https://api.adviceslip.com/advice",
	)

	for res := range ch {
		data, err := getResponseData(res)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("returned data:")
		fmt.Println(data)
	}
}

func getResponseData(resp *http.Response) (string, error) {
	if resp == nil {
		return "", errors.New("no response")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	var buf bytes.Buffer
	err = json.Indent(&buf, body, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal body into json: %v", err)
	}

	return buf.String(), nil
}
