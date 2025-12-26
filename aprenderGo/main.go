package main

//1- package main palabra clave

import (
	"aprenderGo/dividir"
	"aprenderGo/matematica"
	"aprenderGo/mathutil"
	"errors"
	"fmt"
	"math"
)

var nivelPaquete = 56

// anadir paquete go mod init nombre del proyecto
// go mod tidy

//2-paquete import math, mathutil, fmt

func main() {

	despedida := decirAdios()
	fmt.Println(despedida)
	//fmt.Println("Hola con fehejo")

	fmt.Println(nivelPaquete)

	fmt.Println(mathutil.Add(2, 3))

	fmt.Println(mathutil.Resta(4, 3))

	fmt.Println(matematica.Multi(4, 3))

	fmt.Println(dividir.Dividir(24, 2))

	fmt.Println(math.Sqrt(144))

	fmt.Println(mathutil.Edad)
	//3-var
	var x int = 5
	fmt.Println(x)

	var a, b = 3, 3
	fmt.Println(a, b)

	y := 5
	fmt.Println(y)

	Nombre, Edad, Email, Direccion := "fehejo", 56, "fhernandez1@hotmail.com", " Los Restauradores"
	fmt.Println(Nombre, Edad, Email, Direccion)

	fmt.Println(A + B)

	fmt.Println(Pi)

	decirHola()

	suma := add(2, 3)
	fmt.Println(suma)

	divi, err := divi(8, 2)
	if err != nil {
		fmt.Println("hay error")
	}

	fmt.Println(divi)

	if y := 10; y > 5 {
		fmt.Println("y es mayor de 5")

	}

	f := -1
	if f > 0 {
		fmt.Println("es positivo")
	} else if f == 0 {
		fmt.Println("es zero")
	} else {
		fmt.Println("es negativo")
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//range itera sobre un slice
	nums := []int{1, 2, 3}
	for _, num := range nums {
		fmt.Println(num)
	}

}

// 4-cosnt valores que nunca cambia durante la ejecucion del programa
// configuraciones y valores matematicos
const Pi = 3.14

const (
	A = 1
	B = 2
)

// 5-func codigo reutilizable que realizan una tarea especifica

func decirHola() {

	fmt.Println("Hola fehejo")
}
func add(x, y int) int {
	return x + y

}

// 6-if condiciones que se usan cuando son verdaderas

func divi(x, y int) (int, error) {
	if y <= 0 {
		return 0, errors.New("no posee error")
	}

	return x / y, nil

}

//7-else trabaja junto con if para cuando la condicion de if es falsa

//permite ambos casos cuando es verdadero y cuando es falso

//8- for

//9- return de vuelve valor de una funcion

func decirAdios() string {
	return "Adios"

}
