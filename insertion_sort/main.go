// An insertion sort with binary tree
package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func sort(values []int) {
	var tree *tree
	for _, v := range values {
		tree = add(tree, v)
	}
	appendValues(values[:0], tree)
}

func appendValues(values []int, tree *tree) []int {
	return values
}

func add(tree *tree, value int) *tree {
	return nil
}

func main() {
	var values []int
	for i := 100; i > 0; i-- {
		values = append(values, i)
	}
	sort(values)
	fmt.Println(values)
}
