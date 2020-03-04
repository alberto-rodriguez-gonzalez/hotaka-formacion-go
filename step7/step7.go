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

// Order blablabla
type Order struct {
	dishes    [10]Dish
	numDishes int
}

var (
	productList map[string]Product
)

func (o *Order) addDish(name string, quantity int) {
	var amount float32
	var price = productList[name].price

	amount = price * float32(quantity)

	if o.numDishes < 10 {
		dish := Dish{name, quantity, price, amount}
		o.dishes[o.numDishes] = dish
		o.numDishes++
	}
}

func (o *Order) show() {
	for i := 0; i < o.numDishes; i++ {
		fmt.Println(o.dishes[i])
	}
}

func createProductList() {
	productList = make(map[string]Product)
	productList["Patatas Fritas"] = Product{"Patatas Fritas", 3.5}
	productList["Ensalada Cesar"] = Product{"Ensalada Cesar", 5.5}
}

func main() {
	createProductList()
	var order *Order = new(Order)

	order.addDish("Patatas Fritas", 3)
	order.addDish("Ensalada Cesar", 2)

	order.show()

	newOrder := Order{}
	newOrder.addDish("Ensalada Cesar", 2)
	newOrder.show()

}
