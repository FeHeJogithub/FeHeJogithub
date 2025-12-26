package main

import "fmt"

//trabaja con referencia a ubicaciones de memoria
//copia valores, cuando y como usar pointer

//direccion de la memoria
type LargeStruct struct {
	Data [1000]int
	Name string
	ID   int
}

func main() {

	large := LargeStruct{

		Name: "Mi esturctura grande",
		ID:   1,
	}
	ProcessByValue(large)
	ProcessByPointer(&large)

	x := 10

	p := &x //direccion de la memoria

	fmt.Println(p)

	fmt.Println(*p)

	x1 := 20
	p1 := &x1

	*p1 = 50

	*p1 = 500 //copia de la memoria
	fmt.Println(x1, p1, *p1)

	num := 10

	incrementarPorValor(num)
	fmt.Println("Despues de incrementPorValor:", num)

	incrementarPorPointer(&num)
	fmt.Println("Despues de incrementPorValor:", num)

}

func incrementarPorValor(x int) {

	x = x + 1
	fmt.Println("Dentro de la funcion:", x)

}
func incrementarPorPointer(x *int) {
	*x = *x + 1

	fmt.Println("Dentro de la funcion:", *x)

}

func ProcessByValue(ls LargeStruct) {
	fmt.Println("Procesando: ", ls.Name, ls.ID)

}
func ProcessByPointer(ls *LargeStruct) {
	fmt.Println("Procesando: ", ls.Name, ls.ID, ls.Data)

}
