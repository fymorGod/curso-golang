package main

import (
	"fmt"
	"time"
)

//Channel(canal) - é forma de comunicação entre goroutines
// channel é um tipo assim como bolean ou inteiro

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)

	c <- 2 * base // enviando dados para o canal
	time.Sleep(time.Second)

	c <- 3 * base
	time.Sleep(time.Second)

	c <- 4 * base
	time.Sleep(time.Second)

	c <- 5 * base
	time.Sleep(time.Second)
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")
	a, b := <-c, <-c //
	fmt.Println("B")
	fmt.Println(a, b)
	fmt.Println(<-c)
}
