package model

import (
	"context"

	"github.com/google/uuid"
)

type SensorRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*Sensor, error)
	Create(ctx context.Context, sensor *Sensor) (uuid.UUID, error)
}

type SensorEventService interface {
	GetSensorEvents(ctx context.Context, id_sensor uuid.UUID) (*Sensor, []*Event, error)
}

type SensorService interface {
	Get(ctx context.Context, id uuid.UUID) (*Sensor, error)
	Create(ctx context.Context, sensor *Sensor) (uuid.UUID, error)
}
type Sensor struct {
	ID         uuid.UUID  `json:"id"`
	Nome       string     `json:"nome"`
	Nomeregiao EnumRegiao `json:"nomeregiao"`
	Nomepais   EnumPais   `json:"nomepais"`
}

func NewSensor(nome string, pais EnumPais, regiao EnumRegiao) (Sensor, error) {
	if nome == "" {
		return Sensor{}, ErrInvalidNomeSensor
	}

	if pais == 0 || regiao == 0 {
		return Sensor{}, ErrInvalidLocSensor
	}

	return Sensor{
		Nome:       nome,
		Nomeregiao: regiao,
		Nomepais:   pais,
	}, nil

}

func (sensor *Sensor) Validate(stage EnumStages) error {

	if sensor.Nome == "" {
		return ErrInvalidNomeSensor
	}

	if sensor.Nomepais == 0 || sensor.Nomeregiao == 0 {
		return ErrInvalidLocSensor
	}

	if stage == Get && sensor.ID == uuid.Nil {
		return ErrRequiredIDSensor
	}

	return nil
}
