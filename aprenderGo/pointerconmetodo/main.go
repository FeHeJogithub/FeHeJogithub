package main

import "fmt"

type CuentaDeBanco struct {
	Dueño   string
	Balance float64
}

func (ba CuentaDeBanco) DepositoPorValor(amount float64) {
	ba.Balance += amount
	fmt.Printf("Saldo despues del deposito(metodo por valor): %.2f\n", ba.Balance)
}

func (ba *CuentaDeBanco) Deposito(cantidad float64) {
	ba.Balance += cantidad
	fmt.Printf("Saldo despues del deposito(metodo por valor): %.2f\n", ba.Balance)
}

func (ba *CuentaDeBanco) Retirar(cantidad float64) {
	if ba.Balance >= cantidad {
		ba.Balance -= cantidad
		fmt.Printf("Retiro exitoso. Saldo despues del deposito(metodo por valor): %.2f\n", ba.Balance)
	} else {
		fmt.Println("Fondos insuficientes")
	}

}
func main() {

	account := CuentaDeBanco{
		Dueño:   "Fehejo",
		Balance: 1000.00,
	}

	fmt.Printf("Saldo inicial: %.2f\n", account.Balance)

	account.DepositoPorValor(100.00)
	fmt.Printf("Saldo real despues de deposito By value: %.2f\n", account.Balance)
	account.Deposito(200.00)
	fmt.Printf("Saldo real despues de deposito: %.2f\n", account.Balance)
	account.Retirar(150.00)
	fmt.Printf("Saldo FInal: %.2f\n", account.Balance)

}
