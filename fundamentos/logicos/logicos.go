package main

import "fmt"

//três tipos de retorno
func comprar(trabalho_1, trabalho_2 bool) (bool, bool, bool) {
	comprarTv50 := trabalho_1 && trabalho_2
	comprarTv32 := trabalho_1 != trabalho_2 //ou exclusivo
	comprarSorvete := trabalho_1 || trabalho_2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
