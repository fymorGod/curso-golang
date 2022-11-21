package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	if numeros != nil {
		return total / float64(len(numeros))
	} else {
		return 0.0
	}
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9))
}
