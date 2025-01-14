package main

import (
	"fmt"
	"math/rand"
)

func main() {
	chanUint := make(chan uint8)
	chanCube := make(chan float64)

	go func() {
		for i := 0; i < 10; i++ {
			chanUint <- uint8(rand.Intn(10))
		}
		close(chanUint)
	}()

	go func() {
		for x := range chanUint {
			y := float64(int(x) * int(x) * int(x))
			chanCube <- y
		}
		close(chanCube)
	}()

	for x := range chanCube {
		fmt.Printf("| %.1f |", x)
	}
	fmt.Println()
}
