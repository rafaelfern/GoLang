package main

import "fmt"

//Imovel tipo de estrutura
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := new(Imovel)
	//%p e &casa serve para retornar o endereço de memória onde a casa está alocada.
	fmt.Printf("\ncasa é: %p - %+v\n", &casa, casa)

	chacara := Imovel{17, 28, "chacara linda", 280000}
	fmt.Printf("Chácara é: %p - %+v\n", &chacara, chacara)

	mudaImovel(&chacara)
	fmt.Printf("Chácara é: %p - %+v\n\n", &chacara, chacara)

}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5

}
