package service

import (
	"context"
	"sensor/internal/sensor/repository/model"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

//DI IC - não dizemos qual repo e sim uma interface, dando abertura para mudanças
//de banco de dados sem grandes refatoração
type EventService struct {
	repository model.EventRepository
	logger     *zap.Logger
}

type EventConfig struct {
	RepositoryEvent model.EventRepository
	logger          *zap.Logger
}

func NewEventService(config *EventConfig) *EventService {
	return &EventService{config.RepositoryEvent, config.logger}
}

func (s *EventService) Get(ctx context.Context, id uuid.UUID) (*model.Event, error) {
	event, err := s.repository.Get(ctx, id)
	if err != nil {
		return &model.Event{}, err
	}
	return event, nil

}

func (s *EventService) Create(ctx context.Context, evento *model.Event) (uuid.UUID, error) {
	id_event, err := s.repository.Create(ctx, evento)
	if err != nil {
		return uuid.Nil, err
	}
	return id_event, nil

}
