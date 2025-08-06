package main

import "fmt"

func Sumar(a, b int) int {

	return a + b

}

func Mayor(a, b int) int {
	if a > b {
		return a

	}
	return b
}

func Fibonacci(n int) int {

	if n <= 1 {
		return n

	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	resultado := Sumar(3, 5)
	fmt.Println("Resultado:", resultado) // Output: Resultado: 8

	fmt.Println(Mayor(10, 5)) // Resultado: 10
	fmt.Println(Mayor(7, 9))  // Resultado: 9
	fmt.Println(Mayor(4, 4))  //

	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, Fibonacci(i))

	}

}
