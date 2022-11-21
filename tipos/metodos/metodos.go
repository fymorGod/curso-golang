package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastName string
}

func (p person) getNameComplete() string {
	return p.name + " " + p.lastName
}

func (p *person) setNameComplete(nameComplete string) {
	parties := strings.Split(nameComplete, " ")
	p.name = parties[0]
	p.lastName = parties[1]
}

func main() {
	p1 := person{name: "Pedro", lastName: "Silva"}
	fmt.Println(p1.getNameComplete())

	p1.setNameComplete("Maria Pereira")
	fmt.Println(p1.getNameComplete())
}
