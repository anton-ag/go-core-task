package main

import "fmt"

func Compare(s1, s2 []int) (bool, []int) {
	var result []int

OUTER:
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				result = append(result, v1)
				continue OUTER
			}
		}
	}

	if len(result) == 0 {
		return false, result
	}
	return true, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	_, res := Compare(a, b)
	fmt.Printf("Comparison: %v\n", res)
}
