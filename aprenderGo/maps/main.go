package main

import "fmt"

func main() {
	//no tiene orden
	//sintaxis de map
	m := make(map[string]int)
	m["apple"] = 256

	m["banana"] = 3
	fmt.Println(m["apple"])
	fmt.Println(m["banana"])
}
