package main

import "testing"

func recursiveFibonacci(n uint64) uint64 {
	if n < 2 {
		return n
	}
	return recursiveFibonacci(n-2) + recursiveFibonacci(n-1)
}

func TestFibonacci(t *testing.T) {
	for n := uint64(0); n < 40; n++ {
		if got, want := fibonacci(n), recursiveFibonacci(n); got != want {
			t.Fatalf("fibonacci(%d): %d != %d", n, got, want)
		}
	}
}
