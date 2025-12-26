package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i) //0,1,2

	}

	x := 2

	switch x {

	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")
		break //es redundante pero continua

	case 3:
		fmt.Println("Three")

	}
	fmt.Println("Switch complete")
}
