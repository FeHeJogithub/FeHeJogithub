package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}
	for i, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Printf("Index: %s, Value: %d\n", key, value)
	}
}
