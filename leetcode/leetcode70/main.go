package main

import (
	"fmt"
)

// Función que calcula cuántas formas hay de subir n escalones
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2 // f(1)=1, f(2)=2

	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func main() {
	var n int
	//var i, a, b int
	fmt.Print("Ingrese el número de escalones: ")
	fmt.Scan(&n)

	ways := climbStairs(n)
	fmt.Printf("Número de formas de subir %d escalones: %d\n", n, ways)
	//fmt.Println("variable: ", i, a, b)
}
