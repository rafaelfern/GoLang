package main

import (
	"encoding/json"
	"exercicios/json/model"
	"fmt"
)

//Response1 é estrutura 1
type Response1 struct {
	Page   int
	Fruits []string
}

//Response2 é estrutura 2
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

//ColorGroup é um exemplo de golang.org sobre JSON
type ColorGroup struct {
	ID     int `json:"identificador"`
	Name   string
	Colors []string
}

func main() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"crimson", "red", "ruby", "Maroon"},
	}

	b, _ := json.Marshal(group)
	fmt.Println("JSON: ", string(b))

	m := model.Message{"Alice", "Hello", 123456789}
	novoJSON, _ := json.Marshal(m)
	fmt.Println("Messsage: ", string(novoJSON))

	//Para decode o JSON
	var m1 model.Message

	err := json.Unmarshal(novoJSON, &m1)
	fmt.Println("JSON decodificado: ", m1)
}
