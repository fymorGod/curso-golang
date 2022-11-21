package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      11234.32,
		"Gabriela Silva": 14000.50,
		"Luiz Fylip":     30000.50,
	}
	funcsESalarios["Rafael filho"] = 1756.80
	delete(funcsESalarios, "inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
