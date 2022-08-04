package service

import (
	"context"
	"sensor/internal/sensor/repository/model"

	"github.com/google/uuid"
)

//DI IC - não dizemos qual repo e sim uma interface, dando abertura para mudanças
//de banco de dados sem grandes refatoração
// EventoSensorService fica todas as dependências para esse serviço
type EventSensorService struct {
	repositorySensor model.SensorRepository
	repositoryEvent  model.EventRepository
}

// EventoSensorConfig responsável pela injetar todas as interfaces necessárias no EventoSensorService
type EventSensorConfig struct {
	RepositorySensor model.SensorRepository
	RepositoryEvent  model.EventRepository
}

// NewEventoSensorService
func NewEventoSensorService(config *EventSensorConfig) *EventSensorService {
	return &EventSensorService{
		repositorySensor: config.RepositorySensor,
		repositoryEvent:  config.RepositoryEvent}
}

func (s *EventSensorService) GetSensorEvents(ctx context.Context, id_sensor uuid.UUID) (*model.Sensor, []*model.Event, error) {
	sensor, err := s.repositorySensor.Get(ctx, id_sensor)
	if err != nil {
		return &model.Sensor{}, []*model.Event{}, err
	}
	events, err := s.repositoryEvent.GetEventsToIDSensor(ctx, id_sensor)
	if err != nil {
		return &model.Sensor{}, []*model.Event{}, err
	}

	return sensor, events, nil

}
