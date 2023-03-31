package main

import "testing"

func testFactorial(n uint64) uint64 {
	if n < 2 {
		return 1
	}
	return testFactorial(n-1) * n
}

func TestFactorial(t *testing.T) {
	for n := uint64(0); n <= 20; n++ {
		if got, want := factorial(n), testFactorial(n); got != want {
			t.Fatalf("factorial(%d), %d != %d", n, got, want)
		}
	}
}
