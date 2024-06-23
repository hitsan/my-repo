package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(2)
	if result != "even" {
		t.Errorf("expected:even, actual:%s", result)
	}
}