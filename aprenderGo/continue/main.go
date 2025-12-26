package main

import "fmt"

func main() {

	//omitir ciertos elementos y continuar procesando
	//continue no sale sino que brinca
	for i := 0; i < 5; i++ {

		if i%2 == 0 {

			continue
		}
		fmt.Println(i)
	}

	for i := 0; i <= 3; i++ {
		for j := 1; j <= 3; j++ {

			if j == 2 {

				continue
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}
