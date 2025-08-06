package main

import "testing"

func TestSumar(t *testing.T) {

	/*	result := Sumar(2, 3)
		esperado := 5

		if result != esperado {
			t.Errorf("Error al sumar : se esperaba %d, pero se obtuvo %d", esperado, result)
		}*/

	casos := []struct {
		a, b, esperado int
	}{

		{3, 2, 5},
		{4, 5, 9},
		{456, 123, 579},
		{0, 0, 0},
		{-1, 1, 0},
	}
	for _, caso := range casos {

		resultado := Sumar(caso.a, caso.b)

		if resultado != caso.esperado {
			t.Errorf("Error al sumar : se esperaba %d, pero se obtuvo %d", caso.esperado, resultado)

		}
	}

}
func TestMayor(t *testing.T) {
	casos := []struct {
		a, b, e int
	}{
		{3, 2, 3},
		{4, 3, 4},
		{10, 20, 20},
		{-10, -5, -5},
	}

	for _, caso := range casos {
		r := Mayor(caso.a, caso.b)
		if r != caso.e {
			t.Errorf("Error al comparar: se esperaba %d, pero se obtuvo %d", caso.e, r)
		}
	}
}

//verificar si cubre todo el codigo> go test -cover

//que parte del codigo no esta aprobado 100%> go test  --coverprofile=coverange.out

//go tool cover --html=coverange.out
//color verde esta cubierto color rojo no.

//go test -cover

func TestFibonacci(t *testing.T) {
	casos := []struct {
		n, e int
	}{
		{0, 0},
		{1, 1},
		{7, 13},
		{0, 0},
	}
	for _, caso := range casos {
		r := Fibonacci(caso.n)
		if r != caso.e {
			t.Errorf("Error al calcular Fibonacci: se esperaba %d, pero se obtubo %d", caso.e, r)

		}
	}

}
