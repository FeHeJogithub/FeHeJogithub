package main

//goto transfiere el control a una estiqueta especifica en tu codigo
//se usa en caso de optimizacion

import "fmt"

func main() {

	fmt.Println("Antes de goto")
	goto End
	fmt.Println("Esto no se ejecutara")

End:
	fmt.Println("Despues goto")
}
