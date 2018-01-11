package main

//Para comunicação SINCRONIZADA

import (
	"fmt"
	"time"
)

func pinger(canal chan string) {
	for {
		canal <- "ping"
	}
}

func ponger(canal chan string) {
	for {
		canal <- "pong"
	}
}

func impressora(canal chan string) {
	for {
		msg := <-canal
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	//Canal
	var canal chan string
	canal = make(chan string)
	go pinger(canal)
	go ponger(canal)
	go impressora(canal)

	var entrada string
	//ler entrada do teclado
	fmt.Scanln(&entrada)
}
