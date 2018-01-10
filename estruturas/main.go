package main

import "fmt"

//Imovel é uma struct que armazena dados do imóvel
type Imovel struct {
	X, Y  int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}

	apartamento := Imovel{17, 56, "Meu Ap.", 760000}
	fmt.Printf("\nO apartamento é: %+v\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chácara",
		X:     22,
		valor: 55}

	fmt.Printf("A Chácara é: %+v\n", chacara)

	casa.Nome = "Lar doce Lar"
	casa.valor = 450000
	casa.X = 18
	casa.Y = 29
	fmt.Printf("A casa é: %+v\n\n", casa)
}
