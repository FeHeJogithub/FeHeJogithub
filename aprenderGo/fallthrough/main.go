package main

import "fmt"

func main() {

	x := 3
	switch x {
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:

		fmt.Println("Case 2")
	}

	y := 1
	switch y {
	case 1:
		fmt.Println("Case 1")
		fallthrough

	case 2:
		fmt.Println("Case 2")

	case 3:
		fmt.Println("Case 3")
	}

}
