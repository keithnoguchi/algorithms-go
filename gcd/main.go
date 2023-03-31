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
	inputs := []struct{ x, y int }{
		{0, 0},
		{12, 9},
		{12, 6},
		{10, 0},
		{10, 20},
	}

	for _, arg := range inputs {
		fmt.Printf("gcd(%2d, %2d) = %2d\n", arg.x, arg.y, gcd(arg.x, arg.y))
		fmt.Printf("gcd(%2d, %2d) = %2d\n", arg.y, arg.x, gcd(arg.y, arg.x))
	}
}
