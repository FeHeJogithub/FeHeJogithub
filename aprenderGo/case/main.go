package main

import "fmt"

func main() {
	//defina las diferentes ramas
	//agrupa valores con comportamiento similares
	number := 0

	switch number {

	case 0:
	case 1:
		fmt.Println("number is either o or 1")
	default:
		fmt.Println("Number is something else")
	}
}
