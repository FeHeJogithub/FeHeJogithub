package main

import "fmt"

func main() {

	slice := []int{1, 2, 3}
	for i, j := range slice {

		//slice = append(slice, 4, 5)
		fmt.Println(slice, i, j)

		arr := [5]int{10, 20, 30, 40, 50}
		slice := arr[1:4]
		fmt.Println(slice)

	}

}
