package main

import (
	"fmt"
	"time"
)

// crea hilo separado para ejecutar otras funciones
// hilo separado del principal
func Saluda() {

	fmt.Println("Hola desde un goroutine con fehejo")
}
func main() {

	go Saluda()

	time.Sleep(time.Second)

	fmt.Println("main function completed")

	go printNumbers(5)
	go printNumbers(5)
	time.Sleep(time.Second)
}

func printNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

//concurrencia imprime uno luego otro
