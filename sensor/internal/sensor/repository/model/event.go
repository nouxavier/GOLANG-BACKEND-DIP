package model

import (
	"context"

	"github.com/google/uuid"
)

type EventRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*Event, error)
	GetEventsToIDSensor(ctx context.Context, id_idsensor uuid.UUID) ([]*Event, error)
	Create(ctx context.Context, evento *Event) (uuid.UUID, error)
}

type EventService interface {
	Get(ctx context.Context, id uuid.UUID) (*Event, error)
	Create(ctx context.Context, evento *Event) (uuid.UUID, error)
}
type Event struct {
	ID       uuid.UUID `json:"id"`
	Valor    string    `json:"valor"`
	IDSensor uuid.UUID `json:"idSensor"`
}

func NewEvento(valor string, idSensor string) (Event, error) {
	if valor == "" {
		return Event{}, ErrInvalidValorEvento
	}

	uidSensor, erro := uuid.Parse(valor)
	if erro != nil {
		return Event{}, erro
	}

	return Event{
		Valor:    valor,
		IDSensor: uidSensor,
	}, nil
}
func (evento *Event) Validate(stage EnumStages) error {

	if evento.Valor == "" && stage == Create {
		return ErrRequiredValorEvento
	}

	if evento.IDSensor == uuid.Nil || stage == Create {
		return ErrRequiredValorEvento
	}
	return nil
}
