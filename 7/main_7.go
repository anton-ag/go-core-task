package main

import (
	"fmt"
	"sync"
	"time"
)

func MergeChannels(channels ...chan string) chan string {
	out := make(chan string)
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(channels))

		for _, ch := range channels {
			go func(c chan string) {
				defer wg.Done()
				for msg := range c {
					out <- msg
				}
			}(ch)
		}
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		defer close(c1)

		for i := 0; i < 15; i++ {
			time.Sleep(20 * time.Millisecond)
			c1 <- fmt.Sprintf("A-%v ", i)
		}
	}()

	go func() {
		defer close(c2)

		for i := 0; i < 10; i++ {
			time.Sleep(50 * time.Millisecond)
			c2 <- fmt.Sprintf("B-%v ", i)
		}
	}()

	go func() {
		defer close(c3)

		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			c3 <- fmt.Sprintf("C-%v ", i)
		}
	}()

	out := MergeChannels(c1, c2, c3)

	for msg := range out {
		fmt.Print(msg)
	}
	fmt.Println()
}
