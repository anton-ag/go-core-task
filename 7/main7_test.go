package main

import (
	"fmt"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		defer close(c1)
		for i := 0; i < 3; i++ {
			c1 <- fmt.Sprint("A")
		}
	}()

	go func() {
		defer close(c2)
		for i := 0; i < 4; i++ {
			c2 <- fmt.Sprint("A")
		}
	}()

	out := MergeChannels(c1, c2)

	correctS := "AAAAAAA"
	var s string
	for msg := range out {
		s = s + msg
	}

	if s != correctS {
		t.Errorf("Expected %v, received %v", correctS, s)
	}
}
