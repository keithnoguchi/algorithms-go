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
	fmt.Printf("gcd(20, 10) = %d\n", gcd(20, 10))
	fmt.Printf("gcd(10, 20) = %d\n", gcd(10, 20))
	fmt.Printf("gcd(36, 24) = %d\n", gcd(36, 24))
	fmt.Printf("gcd(24, 36) = %d\n", gcd(36, 24))
}
