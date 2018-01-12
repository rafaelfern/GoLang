package model

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

//Métdodo que vai salvar o conteúdo do BODY em um arquivo txt.
//0600 é permissão para ler e escrever no arquivo
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("[MODEL Page] Erro ao carregar a página. Erro: ", err.Error())
	}
	return &Page{Title: title, Body: body}, nil
}
