package main

import (
	"reflect"
	"testing"
)

var (
	testS1 = []string{"i", "you", "he"}
	testS2 = []string{"i", "me", "he", "him"}
	testS3 = []string{"i", "we", "you", "you", "he", "they"}
)

func TestNotEmptyCompare(t *testing.T) {
	correctResult := []string{"you"}
	result := Compare(testS1, testS2)
	if !reflect.DeepEqual(result, correctResult) {
		t.Errorf("Expected %v, received %v", correctResult, result)
	}
}

func TestEmptyCompare(t *testing.T) {
	correctResult := []string{}
	result := Compare(testS1, testS3)
	if !reflect.DeepEqual(result, correctResult) {
		t.Errorf("Expected %v, received %v", correctResult, result)
	}
}
