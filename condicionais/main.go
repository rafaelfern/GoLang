package main

import (
	"fmt"
)

func main() {
	meses := 0
	situacao := true
	cidade := "teste"

	if meses <= 6 {
		fmt.Println("Esse credor deve a pouco tempo.")
	}

	if situacao {
		fmt.Println("Ele esta devendo")
	} else {
		fmt.Println("Ele esta adimplente")
	}

	if cidade == "teste" {
		fmt.Println("O estado dele é teste")
	}

	if descricao, status := tempoQueDeve(meses); status {
		fmt.Println("Qual a situação do cliente? ", descricao)
		return
	}

	fmt.Println("Obrigado por nos consultar")
}

func tempoQueDeve(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo"
		return
	}
	descricao = "O cleinte está com o carnê em dia."
	return
}
