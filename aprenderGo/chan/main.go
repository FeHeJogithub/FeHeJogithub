package main

import "fmt"

//los canales permiten que la goroutine envien y reciban datos de forma segura

//no comunique compartiendo memoria, comparte memoria comunicando

//go func(){}funcion anonima
func main() {

	/*	ch := make(chan int)
		go func() {
			ch <- 42
		}()
		val := <-ch
		fmt.Println(val) */ // output: 42

	ch := make(chan string, 2)
	ch <- "Hello con fehejo"
	ch <- "World con fehejo"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
