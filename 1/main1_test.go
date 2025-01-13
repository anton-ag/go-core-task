package main

import "testing"

const (
	correctStringify = "125028602.71828testingtrue(1+1i)"
	correctHash      = "bafcecb145dd7638656c53dcd75e13f27329efd1f10cfdec4f9d7135f12ba94e"
)

var (
	i int64     = 12
	o int64     = 062
	d int64     = 0xB2C
	f float64   = 2.71828
	s string    = "testing"
	b bool      = true
	c complex64 = 1 + 1i
)

func TestStringify(t *testing.T) {
	result := Stringify(i, o, d, f, s, b, c)
	if result != correctStringify {
		t.Errorf("Expected %s, received %s", correctStringify, result)
	}
}

func TestHashSHA256(t *testing.T) {
	bSlice := []byte(Stringify(i, o, d, f, s, b, c))
	hash := HashSHA256(bSlice, salt)
	if hash != correctHash {
		t.Errorf("Expected %s, received %s", correctHash, hash)
	}
}

