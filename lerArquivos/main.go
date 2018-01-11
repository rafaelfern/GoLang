package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"exercicios/lerArquivos/model"
	"fmt"
	"os"
	"strings"
)

func main() {

	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. ", err.Error())
		return
	}

	//Antes de encerrar o programa fecha o arquivo.
	defer arquivo.Close()

	/*scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Println("O conteúdo da linha é: ", linha)
	}*/

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler um arquivo com leitor CSV.", err.Error())
		return
	}

	arquivoJSON, err := os.Create("cidades.json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. ", err.Error())
		return
	}

	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString("[\n")

	for _, linha := range conteudo {
		for indiceItem, item := range linha {
			//Elimina determinado caractere de uma sequência de strings
			dados := strings.Split(item, "/")
			//fmt.Printf("Linha[%d] : %s\n", indiceItem, dados)
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]
			fmt.Printf("Cidade: %+v\n", cidade)
			cidadeJSON, err := json.Marshal(cidade)
			if err != nil {
				fmt.Println("[main] Houve um erro ao gerar o JSON do item ", item, ".Erro: ", err.Error())
			}
			escritor.WriteString("  " + string(cidadeJSON))
			if (indiceItem + 1) < len(linha) {
				escritor.WriteString(",\n")
			}
		}
	}
	escritor.WriteString("\n]")
	escritor.Flush()
	arquivoJSON.Close()

	arquivo.Close()
}
