// A gratest common divisor
package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	var tests = []struct{ x, y int }{
		{0, 0},
		{12, 9},
		{12, 6},
		{10, 0},
		{10, 20},
	};

	for _, t := range tests {
		fmt.Printf("gcd(%2d, %2d) = %2d\n", t.x, t.y, gcd(t.x, t.y))
		fmt.Printf("gcd(%2d, %2d) = %2d\n", t.y, t.x, gcd(t.y, t.x))
	}
}
