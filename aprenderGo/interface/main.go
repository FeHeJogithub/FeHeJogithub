package main

import "fmt"

//interface es un conjunto de metodos
//es un contrato, si el metodo tiene todo los de la interface
type Hablador interface {
	Hablar() string
}

type Perro struct{}

func (p Perro) Hablar() string {
	return "Wuuf"
}

func main() {
	var p Hablador = Perro{}

	fmt.Println(p.Hablar())

	imprimirCulaquierCosa(42)
	imprimirCulaquierCosa("hola")
	imprimirCulaquierCosa(3.14)

}

func imprimirCulaquierCosa(v interface{}) {
	fmt.Println(v)

}
