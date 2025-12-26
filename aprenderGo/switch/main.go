package main

import "fmt"

func main() {

	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Other day")
	}

	letter := "a"
	switch letter {

	case "a", "b", "c":
		fmt.Println("Vocal")
	default:
		fmt.Println("Consonante")
	}

}
