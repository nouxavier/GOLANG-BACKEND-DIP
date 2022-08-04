package service

import (
	"context"
	"fmt"
	"sensor/internal/sensor/repository/model"
	"testing"

	"sensor/internal/sensor/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet_sensorService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockSensorResp := &model.Sensor{
			ID:         uid,
			Nome:       "Sensor 1",
			Nomeregiao: model.EnumCentroOeste,
			Nomepais:   model.EnumBrasil,
		}

		mockSensorRepository := new(mocks.MockSensorRepository)
		s := NewSensorService(&SensorConfig{
			RepositorySensor: mockSensorRepository,
		})
		mockSensorRepository.On("Get", mock.Anything, uid).Return(mockSensorResp, nil)

		ctx := context.TODO()
		u, err := s.Get(ctx, uid)

		assert.NoError(t, err)
		assert.Equal(t, u, mockSensorResp)
		mockSensorRepository.AssertExpectations(t)
	})
	t.Run("Error", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockSensorRepository := new(mocks.MockSensorRepository)
		s := NewSensorService(&SensorConfig{
			RepositorySensor: mockSensorRepository,
		})
		mockSensorRepository.On("Get", mock.Anything, uid).Return(nil, fmt.Errorf("Some error down the call chain"))

		ctx := context.TODO()
		u, err := s.Get(ctx, uid)

		assert.Equal(t, u, &model.Sensor{})
		assert.Error(t, err)
		mockSensorRepository.AssertExpectations(t)
	})
}

func TestCreate_sensorService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		mockSensor := &model.Sensor{
			Nome:       "Sensor 1",
			Nomeregiao: model.EnumCentroOeste,
			Nomepais:   model.EnumBrasil,
		}

		mockSensorRepository := new(mocks.MockSensorRepository)
		s := NewSensorService(&SensorConfig{
			RepositorySensor: mockSensorRepository,
		})

		// We can use Run method to modify the user when the Create method is called.
		//  We can then chain on a Return method to return no error
		mockSensorRepository.
			On("Create", mock.AnythingOfType("*context.emptyCtx"), mockSensor).Return(nil)

		ctx := context.TODO()
		id_sensor, err := s.Create(ctx, mockSensor)

		assert.NoError(t, err)

		assert.NotNil(t, id_sensor)
		mockSensorRepository.AssertExpectations(t)
	})

}
