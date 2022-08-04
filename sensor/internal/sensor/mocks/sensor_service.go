package mocks

import (
	"context"
	"sensor/internal/sensor/repository/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockSensorService struct {
	mock.Mock
}

func (m *MockSensorService) Get(ctx context.Context, id uuid.UUID) (*model.Sensor, error) {
	// args that will be passed to "Return" in the tests, when function
	// is called with a uid. Hence the name "ret"
	ret := m.Called(ctx, id)

	// first value passed to "Return"
	var r0 *model.Sensor
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*model.Sensor)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1

}

func (m *MockSensorService) Create(ctx context.Context, sensor model.Sensor) (uuid.UUID, error) {
	ret := m.Called(ctx, sensor)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return sensor.ID, r0

}
