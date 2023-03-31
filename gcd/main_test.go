package main

import "testing"

func TestGcd(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{10, 0, 10},
		{10, 20, 10},
		{12, 6, 6},
		{12, 9, 3},
	}

	for _, tt := range tests {
		if got, want := gcd(tt.x, tt.y), tt.want; got != want {
			t.Fatalf("%v != %v", got, want)
		}
		if got, want := gcd(tt.y, tt.x), tt.want; got != want {
			t.Fatalf("%v != %v", got, want)
		}
	}
}
