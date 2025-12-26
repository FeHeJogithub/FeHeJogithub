package main

import "fmt"

func main() {
	//map define una extructura de datos que almacena clave, valor

	//[clave string] valor int
	m := map[string]int{"manzana": 2, "banana": 3}
	fmt.Println(m["manzana"])

	for key, value := range m {
		fmt.Printf(" key es: %s; value es: %d", key, value)
	}
}
