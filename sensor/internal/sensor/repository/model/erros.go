package model

import (
	"errors"
)

var (
	// ErrInvalidNomeSensor retorna erro de nome inválido.
	ErrInvalidNomeSensor = errors.New("nome do sensor é inválido")
	// ErrInvalidLocSensor retorna erro de localidade inválido.
	ErrInvalidLocSensor = errors.New("localidade do sensor é inválida")
	// ErrInvalidEventoSensor retorna erro de valor inválido.
	ErrInvalidEventoSensor = errors.New("valor do evento do sensor é inválido")
	// ErrInvalidValorEvento retorna valor evento inálido
	ErrInvalidValorEvento = errors.New("valor do evento do sensor é inválido")
	// ErrInvalidIDEvento retorna id envento inválido
	ErrInvalidIDEvento = errors.New("valor do id do evento é inválido")
	// ErrInvalidIDSensor retorna id sensor inválido
	ErrInvalidIDSensor = errors.New("valor do id do sensor é inválido")
	// ErrRequiredIDSensor retorna id sensor obrigatório
	ErrRequiredIDSensor = errors.New("valor do id do sensor é obrigatório")
	// ErrRequiredValorEvento retorna valor sensor obrigatório
	ErrRequiredValorEvento = errors.New("valor do evento é obrigatório")
)
