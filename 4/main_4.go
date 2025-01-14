package main

import "fmt"

func Compare(s1, s2 []string) []string {
	var result []string

OUTER:
	for _, v1 := range s1 {
		result = append(result, v1)
		for _, v2 := range s2 {
			if v1 == v2 {
				result = result[:len(result)-1]
				continue OUTER
			}
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Printf("Result: %v\n", Compare(slice1, slice2))
}
