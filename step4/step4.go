package main

import ( // <- Importacion de paquetes
	"fmt"
)

const precio float32 = 4.0

var array1 [6]int = [6]int{1, 2, 3, 4, 5, 6} // classical way
var array2 = [6]int{1, 2, 3, 4, 5, 6}        // a less verbose way
var array3 = [...]int{1, 2, 3, 4, 5, 6}      // the compiler will count the array elements by itself

var dishes = [10]string{} //inicializa un array de 10 posiciones
var numDishes = 0

func addDish(name string, quantity int) {
	var total float32
	total = precio * float32(quantity)
	if numDishes < 10 {
		dishes[numDishes] = fmt.Sprintf("%d - %s - %.2f", quantity, name, total)
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

}
