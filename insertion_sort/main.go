// An insertion sort.
package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, value := range values {
		root = add(root, value)
	}
	appendValue(values[:0], root)
}

func add(node *tree, value int) *tree {
	if node == nil {
		node = new(tree)
		node.value = value
	} else if value < node.value {
		node.left = add(node.left, value)
	} else {
		node.right = add(node.right, value)
	}
	return node
}

func appendValue(values []int, node *tree) []int {
	if node != nil {
		values = appendValue(values, node.left)
		values = append(values, node.value)
		values = appendValue(values, node.right)
	}
	return values
}

func main() {
	var values []int
	for i := 100_000; i >= 0; i-- {
		values = append(values, i)
	}
	Sort(values)
	fmt.Printf("%q\n", values[len(values)-1])
}
