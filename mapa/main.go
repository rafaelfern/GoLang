package main

import "exercicios/mapa/model"

import (
	"fmt"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 22
	casa.Y = 10
	casa.SetValor(600000)
	cache["Casa Amarela"] = casa

	apto := model.Imovel{}
	apto.Nome = "Apto Azul"
	apto.X = 20
	apto.Y = 32
	apto.SetValor(700000)
	cache[apto.Nome] = apto

	fmt.Println("\nHá uma casa amarela no cache?\n")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, o que achei foi: %+v\n", imovel)
	}

	fmt.Println("Qual o tamanho do cache?", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("Qual o tamanho do cache?", len(cache))

}
