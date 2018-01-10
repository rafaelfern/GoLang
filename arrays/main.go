package main

import "fmt"

func main() {

	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 1

	fmt.Println("Qual a capacidade desse array? ", len(teste))

	cantores := [2]string{"Michael Jackson", "John Lennon"}
	fmt.Printf("O que há nesse array?\n %v\n", cantores)

	capitais := [...]string{"Lisboa", "Luanda", "Maputo", "Brasília"}
	fmt.Println("Qual a capacidade desse novo array? ", len(capitais))

	for incide, cidade := range capitais {
		fmt.Printf("Capital[%d] é %s\n", incide, cidade)
	}

}
