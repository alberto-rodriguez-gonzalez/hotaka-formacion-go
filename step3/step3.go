package main

import ( // <- Importacion de paquetes
	"fmt"
)

const precio float32 = 4.0

func addDish(name string, quantity int) string {
	var total float32
	total = precio * float32(quantity)
	return fmt.Sprintf("%d - %s - %.2f", quantity, name, total)
}

func main() {
	//var dishName string = "Patatas Fritas"
	//var quantity int = 3

	//var dishName = "Patatas Fritas"
	//var quantity = 3
	// dishName = 342

	dishName := "Patatas Fritas"
	quantity := 3

	dish := addDish(dishName, quantity)
	fmt.Println(dish)
}
