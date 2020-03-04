package main

import ( // <- Importacion de paquetes
	"fmt"
)

// Product blablabla
type Product struct {
	name  string
	price float32
}

// Dish blablabla
type Dish struct {
	name     string
	quantity int
	price    float32
	amount   float32
}

var (
	productList map[string]Product
	dishes      = [10]Dish{}
	numDishes   = 0
)

func addDish(name string, quantity int) {
	var amount float32
	var price = productList[name].price

	amount = price * float32(quantity)

	if numDishes < 10 {
		dish := Dish{name, quantity, price, amount}
		dishes[numDishes] = dish
		numDishes++
	}
}

func createProductList() {
	productList = make(map[string]Product)
	productList["Patatas Fritas"] = Product{"Patatas Fritas", 3.5}
	productList["Ensalada Cesar"] = Product{"Ensalada Cesar", 5.5}
}

func main() {
	createProductList()

	dishName := "Patatas Fritas"
	quantity := 3

	addDish(dishName, quantity)
	addDish("Ensalada Cesar", 2)

	for i := 0; i < numDishes; i++ {
		fmt.Println(dishes[i])
	}
}
