package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 2
	fmt.Print("O número ", numero, " se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")
	}

	fmt.Println("VocÊ é da família do UNIX?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fallthrough
	case "windows":
		fmt.Println("Usando o windows")
	default:
		fmt.Println("Deixa pra la...")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Ok. Vc pode dormir até mais tarde.")
	default:
		fmt.Println("Vamos la... é dia de trabalho.")
	}

	numero = 10
	fmt.Println("Esse número cabe em um dígito?")
	switch {
	case numero < 10:
		fmt.Println("sim")
	case numero >= 10 && numero < 100:
		fmt.Println("Serve dois dígitos...?")
	case numero >= 100:
		fmt.Println("O número é muito grande")
	}

}
