package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

//Método: função com receiver (receptor)

func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var product1 product
	product1 = product{
		name:     "Lapis",
		price:    1.79,
		discount: 0.05,
	}
	fmt.Println(product1, product1.priceWithDiscount())

	product2 := product{"Notebook", 1989.90, 0.10}
	fmt.Println(product2.priceWithDiscount())
}
