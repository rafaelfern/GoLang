package main

import (
	"exercicios/writingWebApp/model"
	"fmt"
)

func main() {

	p1 := model.Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.Save()
	p2, err := model.LoadPage("TestPage")
	if err != nil {
		fmt.Println("[main] Houve um erro ao carregar a página. Erro: ", err.Error())
		return
	}
	fmt.Println(string(p2.Body))

}
