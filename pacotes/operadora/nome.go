package operadora

import (
	"exercicios/pacotes/prefixo"
	"strconv"
)

//NomeOperadora representa o nome
var NomeOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora recebe uma string, usa-se o strconv.Itoa() para converter int para string
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " - " + NomeOperadora
