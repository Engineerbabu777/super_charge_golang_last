package main

import (
	"fmt"
	"testing"
)

func TestCalculateValues(t *testing.T) {
	// t.Fail()
	var (
		expected = 201
		a        = 10
		b        = 10
	)

	isOk := CalculateValues(a, b)

	if isOk != expected {
		t.Errorf("expected %d but got %d", expected, isOk)
	}

	fmt.Println("Hello from test")
}

// in the current dir run TestCalculateValues func
// go test ./ -v -run TestCalculateValues
