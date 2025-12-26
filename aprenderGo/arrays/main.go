package main

import "fmt"

func main() {

	arr := [3]string{"Go", "Rust", "Python"}

	for i, v := range arr {
		fmt.Printf("Index %d: %s\n", i, v)
	}

}
