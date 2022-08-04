package mocks

import (
	"context"
	"sensor/internal/sensor/repository/model"

	"github.com/stretchr/testify/mock"

	"github.com/google/uuid"
)

type MockSensorRepository struct {
	mock.Mock
}

func (m *MockSensorRepository) Get(ctx context.Context, id uuid.UUID) (*model.Sensor, error) {
	ret := m.Called(ctx, id)

	var r0 *model.Sensor
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.Sensor)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1

}

func (m *MockSensorRepository) Create(ctx context.Context, sensor *model.Sensor) (uuid.UUID, error) {
	ret := m.Called(ctx, sensor)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return uuid.New(), r0

}
