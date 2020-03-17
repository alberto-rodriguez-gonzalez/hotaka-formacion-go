package tests

import (
	"hotaka-formacion-go/step9/restaurante"
	"testing"
)

func TestCreateProductList(t *testing.T) {
	restaurante.CreateProductList()
	if len(restaurante.ProductList) == 0 {
		t.Errorf("Lista de productos vacia")
	}
}

func TestAddDish(t *testing.T) {

}
