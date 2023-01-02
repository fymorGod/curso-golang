package main

import "fmt"

//Estudos de Threads
func main() {
	ch := make(chan int, 1)

	ch <- 1 //enviando o valor 1 para o canal - enviando dados(escrita)
	<-ch    //recebendo dados do canal (leitura)
	ch <- 2
	fmt.Println(<-ch)
}
