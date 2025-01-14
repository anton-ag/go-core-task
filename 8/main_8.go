package main

import "fmt"

type WaitGroup chan struct{}

func New(n int) WaitGroup {
	return make(WaitGroup, n)
}

func (wg WaitGroup) Add(k int) {
	for i := 0; i < k; i++ {
		wg <- struct{}{}
	}
}

func (wg WaitGroup) Remove(k int) {
	for i := 0; i < k; i++ {
		<-wg
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg := New(len(numbers))

	for _, num := range numbers {
		go func(n int) {
			fmt.Printf("| %v |", n*n)
			wg.Add(1)
		}(num)
	}

	wg.Remove(len(numbers))
	fmt.Println()
}
