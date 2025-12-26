package main

import "fmt"

//define tipo basado en tipo existente, anade atracciones
//se comporta como un int,
//son metodos
type MiInt int

type Celsius float64

//metodo es una variable(c Celsius) mas el type
func (c Celsius) ToFahrenheit() float64 {
	return float64(c * 9 / 5 * 32)
}

func main() {
	var m MiInt = 5
	fmt.Println(m)

	temp := Celsius(25)
	fmt.Println(temp.ToFahrenheit())

}
