package main

import "fmt"

type Pessoa struct{
	Nome string;
	Idade int;
}

func (p Pessoa) mostraPessoa() string{
	return fmt.Sprintf("Nome : %s \nIdade : %d", p.Nome, p.Idade);

}

func main(){
	rafael := Pessoa{Nome:"Rafael Augusto", Idade:30};
	
	fmt.Println(rafael.mostraPessoa());
}