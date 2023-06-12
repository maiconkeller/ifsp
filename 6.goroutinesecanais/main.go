package main

import (
	"fmt"
	"time"
)

func imprimirMensagem(texto string, c chan string) {
	time.Sleep(2 * time.Second)
	c <- texto
}

func main() {
	canal := make(chan string)
	go imprimirMensagem("Olá", canal)
	go imprimirMensagem("Mundo!", canal)
	mensagem1 := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem1, mensagem2)
}
