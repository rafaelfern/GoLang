package main

import (
	"encoding/json"
	"exercicios/estruturas_avancado/model"
	"fmt"
)

/*
type Imovel struct {
	X     int `json:"coordenadaaaaaaa_X"`
	Y     int
	Nome  string
	Valor int
}*/

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 22
	casa.Y = 10
	casa.SetValor(600000)

	fmt.Printf("\nInformações sobre a casa: %+v\n", casa)
	fmt.Printf("O valor da casa é: %+v\n", casa.GetValor())

	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON é: ", string(objJSON))

	/*b := Imovel{
		X:     20,
		Y:     10,
		Nome:  "casa",
		Valor: 100000,
	}

	objJSON, _ := json.Marshal(b)
	fmt.Println("JSON : ", string(objJSON))
	*/

}
