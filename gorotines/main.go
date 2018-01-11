package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"exercicios/lerArquivos/model"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var orquestrador sync.WaitGroup

func main() {
	//Não é um processo paralelo, é um processo concorrente.
	//No log da pra verificar que o tempo de início da tradução é bem próximo.
	orquestrador.Add(2)
	go traduzirParaJSON("saopaulo")
	go traduzirParaJSON("riodejaneiro")
	orquestrador.Wait()
}

func traduzirParaJSON(nomeArquivo string) {

	fmt.Println(time.Now(), " - Começando a tradução do arquivo: ", nomeArquivo)
	arquivo, err := os.Open(nomeArquivo + ".csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}

	defer arquivo.Close()

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler um arquivo com leitor CSV.", err.Error())
		return
	}

	arquivoJSON, err := os.Create(nomeArquivo + ".json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. ", err.Error())
		return
	}
	defer arquivoJSON.Close()

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
	fmt.Println(time.Now(), " - A tradução do arquivo: ", nomeArquivo, " foi finalizado com sucesso")
	orquestrador.Done()
}
