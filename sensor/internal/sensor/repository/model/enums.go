package model

// int enumeradores dos paises
type EnumPais int

const (
	EnumPaisIndefinidos EnumPais = 0
	EnumBrasil          EnumPais = 1
)

// int enumeradores dos paises
type EnumRegiao int

const (
	EnumRegiaoIndefinido EnumRegiao = 0
	EnumNorte            EnumRegiao = 1
	EnumSul              EnumRegiao = 2
	EnumSuldeste         EnumRegiao = 3
	EnumCentroOeste      EnumRegiao = 4
	EnumNordeste         EnumRegiao = 5
)

type EnumStages int8

const (
	Create EnumStages = 1
	Get    EnumStages = 2
)
