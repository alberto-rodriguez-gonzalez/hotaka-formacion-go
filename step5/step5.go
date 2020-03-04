package main

import ( // <- Importacion de paquetes
	"fmt"
)

const precio float32 = 4.0

// Dish blablabla
type Dish struct {
	name     string
	quantity int
	price    float32
	amount   float32
}

var (
	dishes    = [10]Dish{}
	numDishes = 0
)

func addDish(name string, quantity int) {
	var amount float32
	amount = precio * float32(quantity)
	if numDishes < 10 {
		dish := Dish{name, quantity, precio, amount}
		dishes[numDishes] = dish
		numDishes++
	}
}

func main() {
	dishName := "Patatas Fritas"
	quantity := 3

	addDish(dishName, quantity)

	addDish("Ensalada Cesar", 2)
	fmt.Println(dishes)

	for i := 0; i < numDishes; i++ {
		fmt.Println(dishes[i])
	}

	//Punteros
	dish := Dish{"Magret de Pato", 1, 12.0, 12.0}
	fmt.Println(dish.name)
	pDish := &dish // Puntero de dish
	pDish.name = "Calamares a la romana"
	fmt.Println(dish.name)

	// Inicializaciones
	p := Dish{"Magret de Pato", 1, 12.0, 12.0}         // tiene un tipo Dish
	q := &Dish{"Calamares a la romana", 1, 12.0, 12.0} // tiene un tipo *Dish
	r := Dish{amount: 1}                               // El reso son implicitas
	s := Dish{}                                        // Valores por defecto de las variables

	fmt.Println(p, q, r, s)

	var d *Dish = new(Dish)
	fmt.Println(d)
	d.name, d.price, d.quantity = "Sopa de caracoles", 23, 12
	fmt.Println(d)
}
