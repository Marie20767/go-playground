package fanin

import "sync"

// takes any number of read only string channels and returns a read only string channel
// reads from all the input channels and writes their values into the output channel
// note - needs to wait for all input channels to be read from (this is where WaitGroup comes in)
func New(channels ...<-chan string) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup

	for _, channel := range channels {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for v := range channel {
				output <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
