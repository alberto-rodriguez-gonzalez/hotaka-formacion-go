package main

import ( // <- Importacion de paquetes

	"hotaka-formacion-go/step9/restaurante"
)

func main() {
	restaurante.CreateProductList()
	var order *restaurante.Order = new(restaurante.Order)

	order.AddDish("Patatas Fritas", 3)
	order.AddDish("Ensalada Cesar", 2)

	order.Show()

	newOrder := restaurante.Order{}
	newOrder.AddDish("Ensalada Cesar", 2)
	newOrder.Show()

}
