package model

//Galinha representa uma ave do tipo galinha
type Galinha interface {
	Cacareja() string
}

//Pata representa uma ave do tipo Pato
type Pata interface {
	Quack() string
}

type Ave struct {
	Nome string
}

func (a Ave) Cacareja() string {
	return "Cócóricó..."
}

func (a Ave) Quack() string {
	return "Quack, quack..."
}
