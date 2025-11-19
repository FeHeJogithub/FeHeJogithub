package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2 // evita overflow

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1 // buscar derecha
		} else {
			right = mid - 1 // buscar izquierda
		}
	}

	return -1
}

func main() {
	// Ejemplo
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	result := search(nums, target)

	if result != -1 {
		fmt.Printf("El número %d se encuentra en el índice %d\n", target, result)
	} else {
		fmt.Printf("El número %d no se encuentra en el arreglo\n", target)
	}
}
