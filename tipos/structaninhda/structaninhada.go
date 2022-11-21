package main

import "fmt"

type item struct {
	productId int
	quantity  int
	price     float64
}

type request struct {
	userId int
	itens  []item
}

func (r request) valueTotal() float64 {
	total := 0.0
	for _, item := range r.itens {
		total += item.price * float64(item.quantity)
	}
	return total
}

func main() {
	request := request{
		userId: 1,
		itens: []item{
			item{
				productId: 1,
				quantity:  2,
				price:     12.10,
			},
			item{
				productId: 2,
				quantity:  1,
				price:     23.49,
			},
			item{
				productId: 11,
				quantity:  100,
				price:     3.13,
			},
		},
	}
	fmt.Printf("The total value of the order is: $ %.2f", request.valueTotal())
}
