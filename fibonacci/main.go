// A fibonacci sequence
package main

import "fmt"

func fibonacci(n uint64) uint64 {
	var x, y uint64 = 0, 1
	for i := uint64(0); i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	for n := uint64(0); n < 1_000; n++ {
		fmt.Printf("fibonacci(%d) = %d\n", n, fibonacci(n))
	}
}
