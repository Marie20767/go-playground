package firstresponse

import (
	"net/http"
	"sync"
)

// Given any number of url strings, returns the response of the first responder
func Fetch(urls ...string) *http.Response {
	responses := make(chan *http.Response)
	for _, url := range urls {
		go func() {
			responses <- doRequest(url)
		}()
	}

	return <-responses
}

func FetchAll(urls ...string) <-chan *http.Response {
	responses := make(chan *http.Response)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Go(func() {
			responses <- doRequest(url)
		})
	}

	go func() {
		wg.Wait()
		close(responses)
	}()

	return responses
}

func doRequest(url string) *http.Response {
	req, _ := http.NewRequest(http.MethodGet, url, http.NoBody)
	client := &http.Client{}
	res, _ := client.Do(req)

	return res
}
