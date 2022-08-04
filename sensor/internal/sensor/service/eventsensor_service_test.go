package service

import (
	"context"
	"testing"

	"sensor/internal/sensor/mocks"
	"sensor/internal/sensor/repository/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetSensorEvents_EventSensorService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		uid_sensor, _ := uuid.NewRandom()

		mockSensorResp := &model.Sensor{
			ID:         uid_sensor,
			Nome:       "Sensor 1",
			Nomeregiao: model.EnumCentroOeste,
			Nomepais:   model.EnumBrasil,
		}

		mockEventResp := []*model.Event{{
			IDSensor: uid_sensor,
			Valor:    "20A",
		},
		}

		mockSensorRepository := new(mocks.MockSensorRepository)
		mockEventRepository := new(mocks.MockEventRepository)
		s := NewEventoSensorService(&EventSensorConfig{
			RepositoryEvent:  mockEventRepository,
			RepositorySensor: mockSensorRepository,
		})
		mockSensorRepository.On("Get", mock.Anything, uid_sensor).Return(mockSensorResp, nil)
		mockEventRepository.On("GetEventsToIDSensor", mock.Anything, uid_sensor).Return(mockEventResp, nil)

		ctx := context.TODO()
		sensor, events, err := s.GetSensorEvents(ctx, uid_sensor)

		assert.NoError(t, err)
		assert.Equal(t, sensor, mockSensorResp)
		assert.Equal(t, events, mockEventResp)
	})

}
