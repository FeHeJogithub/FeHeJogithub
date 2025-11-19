package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		// Si es un cierre
		if open, isClosing := pairs[ch]; isClosing {
			// Si la pila está vacía o no coincide → inválido
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			// Pop
			stack = stack[:len(stack)-1]
		} else {
			// Es apertura → push
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}

func main() {
	var s string
	fmt.Print("Ingrese la cadena: ")
	fmt.Scan(&s)

	if isValid(s) {
		fmt.Println("La cadena es válida.")
	} else {
		fmt.Println("La cadena NO es válida.")
	}
}
