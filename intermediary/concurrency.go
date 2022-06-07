package main

import(
	"fmt"
	"time"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Quantidade de CPUS:", runtime.NumCPU())

	wg.Add(2)

	go request_facebook()
	go request_google()

	fmt.Println("Quantidade de goruntimes rodando", runtime.NumGoroutine())

	wg.Wait()

}

func request_facebook(){

	for x := 0; x<20; x++ {
		time.Sleep(200)
		fmt.Println("Request no facebook feita com sucesso")
	}

	wg.Done()

}

func request_google(){

	for x := 0; x<20; x++ {
		time.Sleep(200)
		fmt.Println("Request no google feita com sucesso")
	}

	wg.Done()

}