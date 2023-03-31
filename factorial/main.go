// A factorial sequence
package main

import "fmt"

func factorial(n uint64) uint64 {
	var x uint64 = 1
	for i := uint64(2); i <= n; i++ {
		x *= i
	}
	return x;
}

func main() {
	for n := uint64(0); n <= 20; n++ {
		fmt.Printf("factorial(%d) = %d\n", n, factorial(n))
	}
}
