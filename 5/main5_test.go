package main

import (
	"reflect"
	"testing"
)

var (
	s1 = []int{1, 2, 3, 4, 5}
	s2 = []int{4, 5, 6, 7, 8}
)

func TestCompare(t *testing.T) {
	correctSlice := []int{4, 5}
	correctBool := true

	b, s := Compare(s1, s2)
	if !(b == correctBool) || !reflect.DeepEqual(s, correctSlice) {
		t.Errorf("Expected %v, %v, received %v, %v", correctBool, correctSlice, b, s)
	}

}
