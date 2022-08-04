package mocks

import (
	"context"
	"sensor/internal/sensor/repository/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockEventSensorService struct {
	mock.Mock
}

func (m *MockEventSensorService) GetSensorEvents(ctx context.Context, id_sensor uuid.UUID) (*model.Sensor, []*model.Event, error) {
	ret := m.Called(ctx, id_sensor)

	var r0 *model.Sensor
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.Sensor)
	}
	var r1 []*model.Event
	if ret.Get(1) != nil {
		r1 = ret.Get(1).([]*model.Event)
	}

	var r2 error

	if ret.Get(2) != nil {
		r2 = ret.Get(2).(error)
	}

	return r0, r1, r2

}
