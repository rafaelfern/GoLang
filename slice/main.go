package main

import "fmt"

func main() {

	var nums []int
	fmt.Println(nums, len(nums), cap(nums))

	//Para incializar um slice precisa usar o comando make()
	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))

	//append para adicionar ao slice
	capitais = append(capitais, "Brasília")
	fmt.Println(capitais, len(capitais), cap(capitais))

	cidades := make([]string, 5)
	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Madeira"
	cidades[3] = "Tokyo"
	cidades[4] = "Singapura"

	fmt.Println(cidades, len(cidades), cap(cidades))

	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] é %s\n", indice, cidade)

	}
	//Cidades[3:2] O primeiro indice conta a partir de zero, após o : conta a partir de 1
	capitaisAsia := cidades[3:5]
	fmt.Println(capitaisAsia)

	//Pegar só os dois primeiros ítens do slice
	temp1 := cidades[:2]
	fmt.Println(temp1)

	//Pegar só os dois últimos do slice
	temp2 := cidades[len(cidades)-2:]
	fmt.Println(temp2)

	//Para retirar algum indice do slice
	indiceDoItemARetirar := 2
	//tem que criar um slice temporário que são os itens da cidade até o indice do item a retirar
	temp := cidades[:indiceDoItemARetirar]
	temp = append(temp, cidades[indiceDoItemARetirar+1:]...)
	copy(cidades, temp)
	fmt.Println(cidades)

}
