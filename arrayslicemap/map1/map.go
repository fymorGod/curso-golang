package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// mapas devem ser incializados
	aprovados := make(map[int]string)

	aprovados[12345678944] = "Maria"
	aprovados[12345678942] = "Luiz"
	aprovados[12345678941] = "Joao"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s(CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678941])
	delete(aprovados, 12345678941)
	fmt.Println(aprovados[12345678941])
}
