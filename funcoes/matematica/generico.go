package matematica

//Calculo genérico que executa qualquer tipo de calculo, basta enviar a função desejada.
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Diminui subtrai dois termos
func Diminui(x int, y int) int {
	return x - y
}

//Multiplicador multiplica dois numeros
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor calcula a divisão do a pelo b
func Divisor(a int, b int) (resultado int) {
	resultado = a / b
	return resultado
}
