package main

import "fmt"

//struct es un objetoen java, forma de crear tipo de datos

type Persona struct {
	Nombre string
	Edad   int
}

type Rectangulo struct {
	Width, Height float64
}

// r es el receptor del metodo
//se basa en la composicion
func (r Rectangulo) Area() float64 {

	return r.Width * r.Height
}

func main() {

	rect := Rectangulo{
		Width:  5,
		Height: 5,
	}
	fmt.Println(rect.Area())

	p := Persona{

		Nombre: "fehejo",
		Edad:   56,
	}
	fmt.Println(p)

}
