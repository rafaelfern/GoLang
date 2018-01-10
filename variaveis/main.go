package main

import (
	"fmt"
)

//Variáveis com inicial maiúscula são visíveis por outros pacotes.
//Variáveis com inicial minúscula são vistas somente dentro do pacote.

var (
	//Endereço é publica
	Endereco            string
	Telefone            string
	quantidade, estoque int     // quantidade = 0
	comprou             bool    //comprou = false
	valor               float64 //valor = 0.00
	palavrasEspeciais   rune
)

func main() {

	//fmt.Println("É do tipo: ", reflect.TypeOf(nome))
	//fmt.Println("É do tipo: ", reflect.TypeOf(x))
	teste := "Valor de teste"

	fmt.Printf("endereço: %s\r\n", Endereco)
	fmt.Printf("Quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavrasEspeciais)
	fmt.Printf("O valor de teste é: %v\r\n", teste)
}
