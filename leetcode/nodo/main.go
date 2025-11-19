package main

import "fmt"

/*type Nodo struct {
	Valor   int
	Vecinos []*Nodo
}

func InOrder(n *Nodo) {
	if n != nil {
		InOrder(n.Izq)
		fmt.Println(n.Valor)
		InOrder(n.Der)
	}
}*/

func main() {

	//fmt.Print(InOrder(n.valor))

	a := []int{1, 2, 3}
	b := a
	b[1] = 99
	fmt.Println(a)
	fmt.Println(b)

	arr := [5]int{10, 20, 30, 40, 50}
	s := arr[1:4]

	fmt.Println(arr, s)

}
