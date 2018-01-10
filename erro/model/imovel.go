package model

import "errors"

//Imovel guarda uma estrutura de imoveis
type Imovel struct {
	X      int    `json:"coordenadaa_x"`
	Y      int    `json:"coordenada_y"`
	Nome   string `json:"nome"`
	valor  int
	estado string
}

//Cria um ponteiro para um objeto do tipo Error

var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido para um imóvel.")
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto.")

//SetValor define o valor do imóvel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 5000 {
		err = ErrValorDeImovelInvalido
		return err
	} else if valor > 100000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

//GetValor retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}
