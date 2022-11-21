package main

import "fmt"

type car struct {
	name           string
	velocityActual int
}

type ferrari struct {
	car         // campos anônimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.velocityActual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com Turbo ligado? %v\n", f.name, f.turboLigado)
}
