// A factorial sequence
package main

import "fmt"

func factorial(n uint64) uint64 {
	if n < 2 {
		return 1
	}
	return factorial(n - 1) * n
}

func main() {
	for n := uint64(0); n < 10; n++ {
		fmt.Printf("factorial(%d) = %d\n", n, factorial(n))
	}
}
