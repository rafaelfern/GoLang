package main

import "fmt"

func main() {
	//for i:=0;i<=10;i++ {
	//	fmt.Println(i);
	//}

	//numero := 1;
	//for numero == 1{
	//	fmt.Println("")
	//}

	mapa := map[string]string{"Cidade": "Rio de Janeiro", "Idade": "400"}

	//Esse for ignora todas as chaves do map, trazendo somente o valor dele.
	for _, values := range mapa {
		fmt.Println(values)
	}

	for keys, values := range mapa {
		fmt.Println("Chave: ", keys, " Valores: ", values)
	}

	texto := "Eu adoro escrever programas usando Go"
	for indice, letra := range texto {
		fmt.Printf("Texto[%d] = %q\n", indice, letra)
	}
}
