package main

import("fmt")

var name = "Samuel" // Variavel com escopo nivel pacote, não pode ser declara com operador curto
var idade int // Pode ser declara sem valor com o tipo no escopo de pacote e só poderá ser atribuida um valor em um bloco de codigo

type int2 int // Para criar o seu propio tipo, que tem como base um tipo primitivo

func main(){

	// := conhecido como operador curto é um operador para declarar variavel com tipagem automatica
	// = para atribuir valor a variaveis existentes

	idade = 21

	fmt.Println("Start script")

	fmt.Println(name)
	fmt.Println(idade)

}