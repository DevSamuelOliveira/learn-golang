package main

import "fmt"

func main() {

 const nome, idade = "Samuel Souza de Oliveira", 21

	// o statement defer é uma pilha de execução significa que o comando será executado ao final da função.

	defer fmt.Println(nome)
	fmt.Println(idade)

}