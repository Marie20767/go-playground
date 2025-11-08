package fanin

import "sync"

// takes any number of read only string channels and returns a read only string channel
// reads from all the input channels and writes their values into the output channel
// note - needs to wait for all input channels to be read from (this is where WaitGroup comes in)
func New(channels ...<-chan string) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup

	for _, channel := range channels {
		wg.Go(func() {
			for s := range channel {
				output <- s
			}
		})
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func FanIn(inputs ...<-chan string) <-chan string {
	c := make(chan string)
	
	// Start a goroutine for each input channel
	for _, input := range inputs {
			go func() {
					for msg := range input {
							c <- msg
					}
			}()
	}
	
	return c
}
