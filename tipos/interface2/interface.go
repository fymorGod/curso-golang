package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo         string
	turboLigado    bool
	velocityActual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{
		modelo:         "F40",
		turboLigado:    false,
		velocityActual: 0,
	}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{
		modelo:         "F40",
		turboLigado:    false,
		velocityActual: 0,
	}

	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
