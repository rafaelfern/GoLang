package main

import "exercicios/erro/model"
import "fmt"
import "encoding/json"

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa da Lucy"
	casa.X = 22
	casa.Y = 39

	if err := casa.SetValor(100); err != nil {

		fmt.Println("[main] Houve um erro na atribuição de valor à casa: ", err.Error())

		return
	}

	fmt.Printf("\nCasa é: %+v\n", casa)

	objJSON, err := json.Marshal(casa)
	//Erro é um ponteiro de um objeto que implementa a interface ERROR
	//Se o ponteiro err está apontando para algum objeto do tipo ERROR
	if err != nil {
		fmt.Println("[main] Houve um erro na geração do objeto JSON: ", err.Error())
		return
	}
	fmt.Println("A casa em JSON: ", string(objJSON))
}
