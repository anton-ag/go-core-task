package main

import (
	"fmt"
	"math/rand"
)

func genereateInts(num int, ch chan int) {
	defer close(ch)
	for i := 0; i < num; i++ {
		ch <- rand.Intn(10)
	}
}

func main() {
	ch := make(chan int)

	go genereateInts(10, ch)

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("| %v |", v)
	}
	fmt.Println()

}
