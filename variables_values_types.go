package main

import("fmt")

var name = "Samuel" // Variavel com escopo nivel pacote, não pode ser declara com operador curto

func main(){

	// := conhecido como operador curto é um operador para declarar variavel com tipagem automatica
	// = para atribuir valor a variaveis existentes

	fmt.Println("Start script")

	fmt.Println(name)

}