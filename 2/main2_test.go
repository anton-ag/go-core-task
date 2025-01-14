package main

import (
	"reflect"
	"testing"
)

var testSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func TestSliceExample(t *testing.T) {
	correctResult := []int{2, 4, 6, 8, 10}
	evenSlice := SliceExample(testSlice)
	if !reflect.DeepEqual(correctResult, evenSlice) {
		t.Errorf("Expected %v, received %v", correctResult, evenSlice)
	}
}

func TestAddElements(t *testing.T) {
	correctResult := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 150}
	expandedSlice := AddElements(testSlice, 150)
	if !reflect.DeepEqual(correctResult, expandedSlice) {
		t.Errorf("Expected %v, received %v", correctResult, expandedSlice)
	}
}

func TestNoErrorRemoveElement(t *testing.T) {
	correctResult := []int{1, 2, 3, 4, 5, 6, 7, 9, 10}
	cutSlice, err := RemoveElement(testSlice, 7)
	if err != nil {
		t.Error(err.Error())
	} else if !reflect.DeepEqual(correctResult, cutSlice) {
		t.Errorf("Expected %v, received %v", correctResult, cutSlice)
	}
}

func TestErrorRemoveElement(t *testing.T) {
	_, err := RemoveElement(testSlice, 123)
	if err == nil {
		t.Error("Expected Index Out Of Range Error")
	}
}
