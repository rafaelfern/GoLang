package main

import (
	"exercicios/funcoes/matematica"
	"fmt"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 3, 3)
	fmt.Printf("O resultado da multiplicação é: %d\n", resultado)
	resultado = matematica.Soma(3, 9)
	fmt.Printf("O resultado da soma é: %d\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 20, 2)
	fmt.Printf("O resultado da divisão é: %d\n", resultado)

}
