package main

import (
	"exercicios/pacotes/operadora"
	"exercicios/pacotes/prefixo"
	"fmt"
)

//NomeDoUsuario é o nome do usuario do sistema
var NomeDoUsuario = "Rafael"

func main() {
	fmt.Printf("Nome do usuário: %s\n", NomeDoUsuario)
	fmt.Printf("O prefixo da capital é: %d\n", prefixo.Capital)
	fmt.Printf("O nome da operadora é: %s\n", operadora.NomeOperadora)
	fmt.Printf("%s\n", operadora.PrefixoDaCapitalOperadora)
}
