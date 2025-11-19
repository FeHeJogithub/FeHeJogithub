package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price // Nuevo precio mínimo encontrado
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice // Calcula ganancia
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("💰 Máxima ganancia:", maxProfit(prices))
}

/*Vamos a recorrer uno por uno:

🔍 PASO A PASO
1️⃣ price = 7
minPrice = 7
maxProfit = 0

2️⃣ price = 1

Como 1 < 7:

minPrice = 1
maxProfit = 0

3️⃣ price = 5

Ganancia = 5 - 1 = 4

Como 4 > 0:

maxProfit = 4

4️⃣ price = 3

Ganancia = 3 - 1 = 2

Como 2 < 4, no cambia nada.

minPrice = 1
maxProfit = 4

5️⃣ price = 6

Ganancia = 6 - 1 = 5

Como 5 > 4:

maxProfit = 5

6️⃣ price = 4

Ganancia = 4 - 1 = 3 → No supera 5.

🎯 RESULTADO FINAL
maxProfit = 5


La mejor compra/venta fue:

Comprar a 1

Vender a 6*/
