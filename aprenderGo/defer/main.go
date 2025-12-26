package main

import "fmt"

func main() {

	//defer se usa para limpiar archivo o para liveral memoria
	//defer retrasa la ejecucion de una funcion a punto de la que retorna
	defer fmt.Println("Esto esta impreso ultimo")

	fmt.Println("Esto esta impreso primero")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Loop finished")
}

//primero en entrar ultimo en salir
