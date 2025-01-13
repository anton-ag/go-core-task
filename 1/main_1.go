package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const salt = "go-2024"

var (
	numDecimal     int64     = 256
	numOctal       int64     = 077
	numHexadecimal int64     = 0xB2B
	pi             float64   = 3.14159
	st             string    = "Task1"
	boolean        bool      = false
	complex        complex64 = 1 + 3i
)

func PrintTypes(vs ...interface{}) {
	for _, v := range vs {
		fmt.Printf("%T\n", v)
	}
}

func Stringify(vs ...interface{}) string {
	var s string
	for _, v := range vs {
		s = s + fmt.Sprintf("%v", v)
	}
	return s
}

func HashSHA256(bs []byte, salt string) string {
	hasher := sha256.New()

	saltBytes := []byte(salt)
	bytesSalted := append(bs, saltBytes...)

	hasher.Write(bytesSalted)
	hashSum := hasher.Sum(nil)

	return hex.EncodeToString(hashSum)
}

func main() {
	// print types
	fmt.Println("Variables' types:")
	PrintTypes(numDecimal, numOctal, numHexadecimal, pi, st, boolean, complex)

	// print string
	fmt.Println("\nCompiled string:")
	s := Stringify(numDecimal, numOctal, numHexadecimal, pi, st, boolean, complex)
	fmt.Println(s)

	// rune slice
	fmt.Println("\nConverting to rune slice:")
	rSlice := []rune(s)
	fmt.Println(rSlice)

	// hash sum
	byteSlice := []byte(s)
	hashSum := HashSHA256(byteSlice, salt)
	fmt.Println("\nHash Sum:")
	fmt.Println("As a string: " + hashSum)
	fmt.Printf("As bytes: %v\n", []byte(hashSum))
}

