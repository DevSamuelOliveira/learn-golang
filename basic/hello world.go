package main // Referencia o arquivo principal

import (
	"fmt" // importa as funções de formatação default
)

func main(){ //Rerefencia a função principal

	// _ para quando recebe um valor não vai ser utilizado posteriomente, em go não é possivel ter uma variavel que não está sendo ou será utilizada
	// as funções printf e println tem como retorno (quantidade de bytes utilizada, erros)

	fmt.Printf("Hello world!\n")
	fmt.Println("Hello world!")
	fmt.Printf("Primeira mensagem com printf %d\n", 100)
	fmt.Println("Primeira mensagem com println", 100)

	qt_bytes, _ := fmt.Println("mensagem")

	fmt.Printf("A quantidade de bytes usada anteriomente foi: %d\n", qt_bytes)

}