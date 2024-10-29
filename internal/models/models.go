package models

type InformacoesBase struct {
	MaxId int
	Count int
}

type Tabelas struct {
	Nome            string
	BaseProducao    InformacoesBase
	BaseHomologacao InformacoesBase
	Status          bool
}
