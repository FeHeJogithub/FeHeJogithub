package main

import (
	"fmt"
)

type Nodo struct {
	Valor int
	Next  *Nodo
}

type Lista struct {
	Head *Nodo
}

// Insertar al inicio
func (l *Lista) InsertarInicio(v int) {
	nuevo := &Nodo{Valor: v, Next: l.Head}
	l.Head = nuevo
}

// Insertar al final
func (l *Lista) InsertarFinal(v int) {
	nuevo := &Nodo{Valor: v}

	if l.Head == nil {
		l.Head = nuevo
		return
	}

	actual := l.Head
	for actual.Next != nil {
		actual = actual.Next
	}
	actual.Next = nuevo
}

// Eliminar un nodo por valor
func (l *Lista) Eliminar(v int) {
	if l.Head == nil {
		return
	}

	if l.Head.Valor == v {
		l.Head = l.Head.Next
		return
	}

	actual := l.Head
	for actual.Next != nil && actual.Next.Valor != v {
		actual = actual.Next
	}

	if actual.Next != nil {
		actual.Next = actual.Next.Next
	}
}

// Recorrer e imprimir
func (l *Lista) Imprimir() {
	actual := l.Head
	for actual != nil {
		fmt.Print(actual.Valor, " -> ")
		actual = actual.Next
	}
	fmt.Println("nil")
}

func main() {

	lista := &Lista{}

	fmt.Println("Insertar al inicio:")
	lista.InsertarInicio(10)
	lista.InsertarInicio(50)
	lista.InsertarInicio(20)
	lista.Imprimir()

	fmt.Println("Insertar al final:")
	lista.InsertarFinal(30)
	lista.InsertarFinal(40)
	lista.InsertarFinal(60)
	lista.Imprimir()

	fmt.Println("Eliminar valor 20:")
	lista.Eliminar(20)
	lista.Imprimir()

	fmt.Println("Eliminar valor 40:")
	lista.Eliminar(40)
	lista.Imprimir()
}
