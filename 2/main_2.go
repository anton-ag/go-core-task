package main

import (
	"fmt"
	"math/rand"
)

func OriginateSlice() []int {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, 1+rand.Intn(10))
	}

	return s
}

func SliceExample(in []int) []int {
	var out []int
	for _, elem := range in {
		if elem%2 == 0 {
			out = append(out, elem)
		}
	}

	return out
}

func AddElements(in []int, num int) []int {
	return append(in, num)
}

func CopySlice(in []int) []int {
	return append([]int{}, in...)
}

func RemoveElement(in []int, index int) ([]int, error) {
	if len(in) <= index {
		return nil, fmt.Errorf("Index out of range")
	}

	return append(in[:index], in[index+1:]...), nil
}

func main() {
	originalSlice := OriginateSlice()
	// originalSlice := []int{2, 2, 2, 1, 1}

	fmt.Printf("Orirginal slice: %v\n\n", originalSlice)

	evenSlice := SliceExample(originalSlice)
	fmt.Printf("Even elements: %v\n\n", evenSlice)

	newSlice := AddElements(evenSlice, 11)
	fmt.Printf("Add 11 at the end: %v\n\n", newSlice)

	copySlice := CopySlice(newSlice)
	copySlice[0] = 13
	fmt.Printf("Make sure this is a deep copy:\n%v\n%v\n\n", copySlice, newSlice)

	shortSlice, err := RemoveElement(copySlice, 3)
	if err != nil {
		fmt.Println("Slice too short")
	} else {
		fmt.Printf("Removed 4th element:\n%v\n", shortSlice)
	}
}
